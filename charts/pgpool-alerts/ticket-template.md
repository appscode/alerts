This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern. To troubleshoot and find what to do let's get started with the database phase.

Let's say you have pgpool database in namespace demo.Please refer to [KubeDB Docs](https://kubedb.com/docs/latest/guides/pgpool/) for more about KubeDB.
```bash
Kubectl get pgpool -n <ns>     // will list all the database in a namesapce  
$ kubectl get pgpool -n demo
NAME                TYPE                  VERSION   STATUS   AGE
pgpool-monitoring   kubedb.com/v1alpha2   4.5.0     Ready    2d5h
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

# Phases
## Ready
KubeDB Managed Database phase becomes Ready when every database server is in the cluster and working properly. There's nothing to worry about.

## Provisioning
Usually, the Database Phase is set to `Provisioning` while bootstrapping for the first time. If you find the database phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the pgpool object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- [Get the pgpool object](#get-pgpool)
- [Describe pgpool object](#describe-pgpool)
- Describe the PetSet object: If there are multiple petsets, describe all of them. To do that, first
  [List petsets under a pgpool](#list-petsets-under-a-pgpool). Then [Describe petsets](#describe-petsets) for each of the petsets.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a pgpool](#list-pods-under-a-pgpool). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Operator logs](#operator-logs)

## Critical
Database Phase Critical means some of the database server/pods is not in the cluster or failing synchronization with the database cluster.
The reasons could be some Database left the cluster maybe for a restart or replication errors or unexpected kills.
To resolve this , we need to  find out which servers/pod that is not in the cluster by checking the logs, describing the database object, or maybe querying in the working database server.

You can contact to AppsCode with the following things attached,
- [Get the pgpool object](#get-pgpool)
- [Describe pgpool object](#describe-pgpool)
- Describe the PetSet object: If there are multiple petset, describe all of them. To do that, first
  [List petsets under a pgpool](#list-petsets-under-a-pgpool). Then [Describe petsets](#describe-petsets) for each of the petsets.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a pgpool](#list-pods-under-a-pgpool). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Operator logs](#operator-logs)

## NotReady
Database Phase NotReady means none of the database servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the database server is Killed, Replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from operator and pod containers, describing the pgpool object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

You can contact to AppsCode with the following things attached,
- [Get the pgpool object](#get-pgpool)
- [Describe pgpool object](#describe-pgpool)
- Describe the PetSet object: If there are multiple petsets, describe all of them. To do that, first
  [List petsets under a pgpool](#list-petsets-under-a-pgpool). Then [Describe petsets](#describe-petsets) for each of the petsets.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a pgpool](#list-pods-under-a-pgpool). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Operator logs](#operator-logs)

# How to ?
## Get pgpool
```bash
kubectl get pgpool -n <namespace> -oyaml <pgpool-object-Name> 
```

## Describe pgpool
```bash
kubectl describe pgpool -n <namespace> <pgpool-object-Name> 
```

## List petsets under a pgpool
```bash
kubectl get petset -n <namespace> -l=app.kubernetes.io/component=connection-pooler,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=pgpools.kubedb.com,app.kubernetes.io/instance=<pgpool-object-Name> 
```

## Describe petsets
```bash
kubectl describe petset -n <namespace> <petset-Name> 
```

## Check secret
```bash
# Get the secret name
kubectl get pp -n <namespace> <pgpool-object-Name> -o jsonpath="{.spec.authSecret.name}"
# Check secret's existence
kubectl get secrets -n <namespace> <pgpool-secret-name>
```

## List pods under a pgpool
```bash
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=connection-pooler,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=pgpools.kubedb.com,app.kubernetes.io/instance=<pgpool-object-Name> 
```

## Describe pods
```bash
kubectl describe pods -n <namespace> <pod-Name> 
```

## Get pod logs
```bash
kubectl logs -n <namespace> <pod-name> -c pgpool
```

## Operator logs
```bash
kubectl logs -n <kubedb-ns> <provisioner-pod-name>
kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
```