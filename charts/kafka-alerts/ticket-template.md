This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern. To troubleshoot and find what to do let's get started with the database phase.

Let's say you have kafka database in namespace demo.
```bash
Kubectl get kafka -n <ns>     // will list all the database in a namesapce  
$ kubectl get kafka -n demo
  NAME    VERSION   STATUS   AGE
  mg-sh   4.4.6     Ready    6h51m
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

# Phases
## Ready
KubeDB Managed Database phase becomes Ready when every database server is in the cluster and working properly. There's nothing to worry about.

## Provisioning
Usually, the Database Phase is set to `Provisioning` while bootstrapping for the first time. If you find the database phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the kafka object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- [Get the kafka object](#get-kafka)
- [Describe kafka object](#describe-kafka)
- Describe the StatefulSet object: If there are multiple sts, describe all of them. To do that, first
  [List statefulsets under a kafka](#list-statefulsets-under-a-kafka). Then [Describe statefulsets](#describe-statefulsets) for each of the sts.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a kafka](#list-pods-under-a-kafka). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded kafka, you also need to check the [mongos bootstrap container's logs](#mongos-bootstrap-logs).
- [Operator logs](#operator-logs)

## Critical
Database Phase Critical means some of the database server/pods is not in the cluster or failing synchronization with the database cluster.
The reasons could be some Database left the cluster maybe for a restart or replication errors or unexpected kills.
To resolve this , we need to  find out which servers/pod that is not in the cluster by checking the logs, describing the database object, or maybe querying in the working database server.

You can contact to AppsCode with the following things attached,
- [Get the kafka object](#get-kafka)
- [Describe kafka object](#describe-kafka)
- Describe the StatefulSet object: If there are multiple sts, describe all of them. To do that, first
  [List statefulsets under a kafka](#list-statefulsets-under-a-kafka). Then [Describe statefulsets](#describe-statefulsets) for each of the sts.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a kafka](#list-pods-under-a-kafka). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded kafka, you also need to check the [mongos bootstrap container's logs](#mongos-bootstrap-logs).
- [Operator logs](#operator-logs)

## NotReady
Database Phase NotReady means none of the database servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the database server is Killed, Replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from operator and pod containers, describing the kafka object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

You can contact to AppsCode with the following things attached,
- [Get the kafka object](#get-kafka)
- [Describe kafka object](#describe-kafka)
- Describe the StatefulSet object: If there are multiple sts, describe all of them. To do that, first
  [List statefulsets under a kafka](#list-statefulsets-under-a-kafka). Then [Describe statefulsets](#describe-statefulsets) for each of the sts.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a kafka](#list-pods-under-a-kafka). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).
- [Operator logs](#operator-logs)

# How to ? 
## Get kafka
```bash
kubectl get kafka -n <namespace> -oyaml <kafka-object-Name> 
```

## Describe kafka
```bash
kubectl describe kafka -n <namespace> <kafka-object-Name> 
```

## List statefulsets under a kafka
```bash
kubectl get sts -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=kafkas.kubedb.com,app.kubernetes.io/instance=<kafka-object-Name> 
```

## Describe statefulsets
```bash
kubectl describe sts -n <namespace> <sts-Name> 
```

## Check secret
```bash
# Get the secret name
kubectl get kf -n <namespace> <kafka-object-Name> -o jsonpath="{.spec.authSecret.name}"
# Check secret's existence
kubectl get secrets -n <namespace> <kafka-secret-name>
```

## List pods under a kafka
```bash
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=kafkas.kubedb.com,app.kubernetes.io/instance=<kafka-object-Name> 
```

## Describe pods
```bash
kubectl describe pods -n <namespace> <pod-Name> 
```

## Get pod logs 
```bash
kubectl logs -n <namespace> <pod-name> -c kafka
```

## Mongos bootstrap logs
```bash
# List mongos pods 
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=kafkas.kubedb.com,app.kubernetes.io/instance=<kafka-object-Name>
# Now check logs
kubectl logs -n <ns> <pod-name>
```

## Operator logs
```bash
kubectl logs -n <kubedb-ns> <provisioner-pod-name>
kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
```