## TODOs on singlestore Critical Alerts

### Database Alerts

- #### SinglestoreInstanceDown
  - Describe the `Singlestore` CR, check the reason from conditions and try restarting the pods
  - Contact AppsCode team
- #### SinglestoreServiceDown
  - Describe the `Singlestore` CR and Try restarting all the pods
  - Contact AppsCode team
- #### SinglestoreTooManyConnections
  - Increase singlestore variable `max_connections`
- #### SinglestoreHighThreadsRunning
  - Increasing singlestore variable `max_connections` may help. 
  - Also try tuning singlestore for memory optimization
- #### SinglestoreRestarted
  - Check if `Singlestore` CR is in Ready status
  - Contact AppsCode team if status is not updated.
- #### SinglestoreHighQPS | SinglestoreHighIncomingBytes | SinglestoreHighOutgoingBytes
  - Scale Singlestore using [KubeDB](https://kubedb.com/docs/latest/guides/singlestore/) Scaling OpsRequest


### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `SinglestoreOpsRequest` is ongoing on same database, Wait until it completes.
  - Contact AppsCode team if this persists for more than 30 minutes.



