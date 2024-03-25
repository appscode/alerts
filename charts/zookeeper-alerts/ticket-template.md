This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern. To troubleshoot and find what to do let's get started with the database phase.

> Please refer to [KubeDB Docs](https://kubedb.com/docs/latest/guides/zookeeper/) for more about KubeDB. You can debug ZooKeeperSentinel objects by changing the CRD name zookeeper to zookeepersentinel in the commands.

Let's say you have ZooKeeper database in namespace demo.
```bash
kubectl get zookeeper -n <ns>  // will list all the database in a namesapce     
$ kubectl get zookeeper -n demo 
  NAME               VERSION   STATUS   AGE
  sample-zookeeper   3.7.2     Ready    1h11m
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

**Ready:** KubeDB Managed Database phase becomes Ready when every database server is in the cluster and working properly. There's nothing to worry about.

**Provisioning:** Usually, the Database Phase is set to `Provisioning` while bootstrapping for the first time. If you find the database phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the ZooKeeper object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- Get the ZooKeeper object:
    ```bash
    kubectl get zookeeper -n <ns> <zookeeper-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml
    ```
- Describe the ZooKeeper object:
    ```bash
      kubectl describe zookeeper -n <ns> <zookeeper-object-Name>
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <statefulset-name>
    ```
- Describe the pods: If there are multiple pods describe all of them one by one
    ```bash
       kubectl describe pod -n <ns> <pod-name>
    ```
- Check the stateful and secret is created:
    ```bash
    kubectl get secret -n <ns> | grep <zookeeper-object-name>
    ```
- Pod logs: If there are multiple pods, log all of them one by one.
    ```bash
    kubectl logs -n <ns> <pod-name> -c zookeeper
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```

**Critical:** Database Phase Critical means some database server/pods is not in the cluster or failing synchronization with the database cluster.
The reasons could be some Database left the cluster maybe for a restart or replication errors or unexpected kills.
To resolve this , we need to  find out which servers/pod that is not in the cluster by checking the logs, describing the database object, or maybe querying in the working database server.

You can contact to AppsCode with the following things attached,

- ZooKeeper object:
    ```bash
    kubectl get zookeeper -n <ns> <zookeeper-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml
    ```
- Describe the ZooKeeper object
    ```bash
      kubectl describe zookeeper -n <ns> <zookeeper-object-Name>
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <statefuleset-name>
    ```
- Describe the pods: if there are multiple pods describe all of them on by one.
    ```bash
       kubectl describe pod -n <ns> <pod-name>
    ```
- Pod logs: If there are multiple pods, log all of them one by one. ZooKeeper in Standalone mode and Cluster mode doesn't contain the `zookeeper-coordinator` container
    ```bash
    kubectl logs -n <ns> <pod-name> -c zookeeper
    kubectl logs -n <ns> <pod-name> -c zookeeper-coordinator
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```

**NotReady:** Database Phase NotReady means none of the database servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the database server is Killed, Replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from operator and pod containers, describing the ZooKeeper object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

In that case please contact AppsCode with the following information.

- ZooKeeper object:
    ```bash
    kubectl get zookeeper -n <ns> <zookeeper-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml
    ```
- Describe the ZooKeeper object
    ```bash
      kubectl describe zookeeper -n <ns> <zookeeper-object-Name>
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <statefulset-name>
    ```
- Describe the pods: if there are multiple pods describe all of them on by one
    ```bash
       kubectl describe pod -n <ns> <podName>
    ```
- Pod logs: If there are multiple pods, log all of them one by one. ZooKeeper in Standalone mode and Cluster mode doesn't contain the `zookeeper-coordinator` container
    ```bash
    kubectl logs -n <ns> <pod-name> -c zookeeper
    kubectl logs -n <ns> <pod-name> -c zookeeper-coordinator
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```