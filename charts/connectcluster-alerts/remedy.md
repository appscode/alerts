## TODOs on ConnectCluster Critical Alerts

### Database Alerts

- #### ConnectClusterWorkerDown
    - Describe the `ConnectCluster` CR, check the reason from conditions and try restarting the pods
    - Check worker pod logs
    - Contact AppsCode team
- #### ConnectClusterTooManyConnections
    - Delete unused connectors and tasks
- #### ConnectClusterConnectorCount
    - Remove unused connectors
- #### ConnectClusterCoordinatorRebalanceFailed
    - Check expected worker pod is running
    - Check worker pod logs
    - Check the ConnectCluster Configurations
- #### ConnectClusterTaskErrorTotalRetries or ConnectClusterTaskTotalFailed or ConnectClusterTaskTotalDestroyed
    - Check connector configurations
    - Check worker pod logs
- #### ConnectClusterTaskTotal
    - Remove unused tasks/connectors

### KubeDB Provisioner

- #### KubeDBConnectClusterPhaseNotReady
    - Describe the `ConnectCluster` CR, check the reason from conditions.
    - Check connectcluster worker pod logs
    - Check KubeDB provisioner logs
    - Contact AppsCode team
- #### KubeDBConnectClusterPhaseCritical
    - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
    - Describe the pods that are not Ready and check their events and conditions.
    - Contact AppsCode team if this persists for more than 30 minutes.