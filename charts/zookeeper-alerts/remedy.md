## TODOs on ZooKeeper Critical Alerts

### Database Alerts

- #### ZooKeeperDown
    - Describe the `ZooKeeper` CR, check the reason from conditions and try restarting the pods
    - Contact AppsCode team
- #### ZooKeeperTooManyNodes
  - Describe the `ZooKeeper` CR, check the reason from conditions and try restarting the `replica` pods
  - Contact AppsCode team
- #### ZooKeeperTooBigMemory
  - Contact AppsCode team
- #### ZooKeeperTooManyWatch
  - Describe the `ZooKeeper` CR, check the reason from conditions and try restarting the `replica` pods
  - Contact AppsCode team
- #### ZooKeeperTooManyConnections
    - Increase ZooKeeper connections in config
- #### ZooKeeperTooManyOpenFiles
    - Contact AppsCode team
### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact AppsCode team
- #### AppPhaseCritical
    - check for `zookeeper` container's logs on the node that is not connected with `Primary` node.
    - If some nodes of the ZooKeeper Cluster are not `Up`, Try restarting those nodes one at a time.
    - Contact AppsCode team if this persists for more than 30 minutes.

