This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern. To troubleshoot and find what to do let's get started with the database phase.

Let's say you have MariaDB database in namespace demo.Please refer to [KubeDB Docs](https://kubedb.com/docs/v2023.01.31/guides/mariadb/) for more about KubeDB.
```bash
Kubectl get mariadb -n <ns>     // wlil list all the database in a namesapce  
$ kubectl get mariadb -n demo
  NAME    VERSION   STATUS   AGE
  mariadb   10.5.8    Ready    6h51m
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

**Ready:** KubeDB Managed Database phase becomes Ready when every database server is in the cluster and working properly. There's nothing to worry about.

**Provisioning:** Usually, the Database Phase is set to `Provisioning` while bootstrapping for the first time. If you find the database phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the MariaDB object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- Get the MariaDB object:
    ```bash
    kubectl get mariadb -n <ns> <mariadb-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the MariaDB object:
    ```bash
      kubectl describe mariadb -n <ns> <mariadb-object-Name> 
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <mariadb-object-Name>
    ```
- Describe the pods: If there are multiple pods describe all of them
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Check the stateful and secret is created:
    ```bash
    kubectl get secret -n <ns> | grep <mariadb-object-name>
    ```
- Pod logs: If there are multiple pods, log all of them. stand alone mariadb doesn't contain the `mariadb-coordinator` container
    ```bash
    kubectl logs -n <ns> <pod-name> -c mariadb
    kubectl logs -n <ns> <pod-name> -c mariadb-coordinator  
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```

**Critical:** Database Phase Critical means some of the database server/pods is not in the cluster or failing synchronization with the database cluster.
The reasons could be some Database left the cluster maybe for a restart or replication errors or unexpected kills.
To resolve this , we need to  find out which servers/pod that is not in the cluster by checking the logs, describing the database object, or maybe querying in the working database server.

You can contact to AppsCode with the following things attached,

- MariaDB object:
    ```bash
    kubectl get mariadb -n <ns> <mariadb-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the MariaDB object
    ```bash
      kubectl describe mariadb -n <ns> <mariadb-object-Name> 
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <mariadb-object-Name>
    ```
- Describe the pods: if there are multiple pods describe all of them
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Pod logs:  if there are multiple pods, log all of them. stand alone mariadb doesn't contain the `mariadb-coordinator` container.
    ```bash
    kubectl logs -n <ns> <pod-name> -c mariadb
    kubectl logs -n <ns> <pod-name> -c mariadb-coordinator  
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```

**NotReady:** Database Phase Not Ready means none of the database servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the database server is Killed, Replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from operator and pod containers, describing the MariaDB object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

In that case please contact AppsCode with the following information.

- MariaDB object:
    ```bash
    kubectl get mariadb -n <ns> <mariadb-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the MariaDB object
    ```bash
      kubectl describe mariadb -n <ns> <mariadb-object-Name> 
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <mariadb-object-Name>
    ```
- Describe the pods: if there are multiple pods describe all of them
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Pod logs:  if there are multiple pods, log all of them. stand alone mariadb doesn't contain the `mariadb-coordinator` container.Include exporter log if it's enabled
    ```bash
    kubectl logs -n <ns> <pod-name> -c mariadb
    kubectl logs -n <ns> <pod-name> -c mariadb-coordinator
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```