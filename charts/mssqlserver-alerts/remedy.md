## TODOs on MSSQLServer Critical Alerts

### Database Alerts

- #### MSSQLServerInstanceDown
  - Describe the `MSSQLServer` CR, check the reason from conditions and try restarting the pods
  - Contact AppsCode team
- #### MSSQLServerTooManyConnections
  - Increase MSSQLServer variable `max_connections`
- #### MSSQLServerHighThreadsRunning
  - Increasing MSSQLServer variable `max_connections` may help. 
  - Also try tuning MSSQLServer for memory optimization
- #### MSSQLServerSlowQueries
  - Please check if `shared_buffers` is 10% to 25% of available Memory.
  - Please check if `effective_cache_size` is 75% of available Memory.
  - Check if you are using proper `VACUUM` method for the relevant table.
- #### MSSQLServerRestarted
  - Check if `MSSQLServer` CR is in Ready status
  - Contact AppsCode team if status is not updated.
- #### MSSQLServerHighQPS | MSSQLServerHighIncomingBytes | MSSQLServerHighOutgoingBytes
  - Scale MSSQLServer using [KubeDB](https://kubedb.com/docs/latest/guides/mssqlserver/) Scaling OpsRequest
- #### MSSQLServerReplicationLag
  - Check if network connection is stable inside kubernetes cluster
  - Check for the memory and cpu resource
  - Contact AppsCode team if stays like this for more than 30 minutes
- #### MSSQLServerSplitBrain
  - Contact AppsCode team 
- #### MSSQLServerTooManyLocksAcquired
  - Check for the connected clients
### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `MSSQLServerOpsRequest` is ongoing on same database, Wait until it completes.
  - check for `Pg-Coordinator` logs on the node that is not connected with `Primary` node.
  - If some nodes of the MSSQLServer Cluster are not `Up`, Try restarting those nodes one at a time.
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If any `MSSQLServerSQLOpsRequest` is ongoing on same database, Wait until it completes. For `Version Upgrade` or some other `OpsRequest` It may take longer then expectation.
  - Contact AppsCode team
- #### OpsRequestFailed
  - Describe the OpsRequest and Check the conditions in it
  - Contact AppsCode team

### Stash Alerts
- #### BackupSessionFailed
  - [Describe the BackupSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-backupsession)
  - Check the conditions in the BackupSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the BackupSession
  - [View the Backup Job/Sidecar log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-backup-jobsidecar-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### RestoreSessionFailed
  - [Describe the RestoreSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-restoresession)
  - Check the conditions in the RestoreSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the RestoreSession
  - [View the Restore Job/Init-Container log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-restore-jobinit-container-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### NoBackupSessionForTooLong
  - Check if the BackupConfiguration is not `Paused`
  - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
  - [Describe the BackupConfiguration](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#backupconfiguration-notready)
  - Check the conditions of BackupConfiguration
  - Check the reasons of the `false` conditions (if any)
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### RepositoryCorrupted
  - Check if the `INTEGRITY` of `repository` is true
  - Contact AppsCode team
- #### RepositoryStorageRunningLow
  - Increase the volume size of `repository` backend
  - Update RetentionPolicy of the BackupConfiguration to free up storage
- #### BackupSessionPeriodTooLong | RestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of `repository` is true
  - Check the `MSSQLServer` CRs status
  - Contact AppsCode team

### KubeStash Alerts
- #### MSSQLServerKubeStashBackupSessionFailed
  - Describe the BackupSession
  - Check the conditions in the BackupSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the BackupSession
  - View the Backup Job log
  - Check if the `INTEGRITY` of Repository is `true`
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### MSSQLServerKubeStashRestoreSessionFailed
  - Describe the RestoreSession
  - Check the conditions in the RestoreSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the RestoreSession
  - View the Restore Job log
  - Check if the `INTEGRITY` of Repository is `true`
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### MSSQLServerKubeStashNoBackupSessionForTooLong
  - Check if the BackupConfiguration is not `Paused`
  - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
  - Describe the BackupConfiguration
  - Check the conditions of BackupConfiguration
  - Check the reasons of the `false` conditions (if any)
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### MSSQLServerKubeStashRepositoryCorrupted
  - Check if the `INTEGRITY` of `repository` is `true`
  - Contact AppsCode team
- #### MSSQLServerKubeStashRepositoryStorageRunningLow
  - Increase the volume size of `repository` backend
  - Update RetentionPolicy to free up storage
- #### MSSQLServerKubeStashBackupSessionPeriodTooLong | MSSQLServerKubeStashRestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of `repository` is `true`
  - Check the `MSSQLServer` CRs status
  - Contact AppsCode team


