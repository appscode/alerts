## TODOs on MySQL Critical Alerts

### Database Alerts

- #### MySQLInstanceDown
  - Describe the `MySQL` CR, check the reason from conditions and try restarting the pods
  - Contact AppsCode team
- #### MySQLServiceDown
  - Describe the `MySQL` CR and Try restarting all the pods
  - Contact AppsCode team
- #### MySQLTooManyConnections
  - Increase mysql variable `max_connections`
- #### MySQLHighThreadsRunning
  - Increasing mysql variable `max_connections` may help. 
  - Also try tuning mysql for memory optimization
- #### MySQLSlowQueries
  - Check Slow Query log file here `/var/log/mysql/mysql-slow.log`
- #### MySQLInnoDBLogWaits
  - Reason for alert: [Innodb_log_waits](https://dev.mysql.com/doc/refman/8.0/en/server-status-variables.html#statvar_Innodb_log_waits)
  - Try reconfiguring [innodb_log_buffer_size](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_log_buffer_size)
- #### MySQLRestarted
  - Check if `MySQL` CR is in Ready status
  - Contact AppsCode team if status is not updated.
- #### MySQLHighQPS | MySQLHighIncomingBytes | MySQLHighOutgoingBytes
  - Scale MySQL using [KubeDB](https://kubedb.com/docs/latest/guides/mysql/) Scaling OpsRequest
- #### MySQLTooManyOpenFiles
  - Increase mysql variable `open_files_limit`

### Group Replication Alerts

- #### MySQLHighReplicationDelay | MySQLHighReplicationTransportTime | MySQLHighReplicationApplyTime | MySQLReplicationHighTransactionTime
  - Try scale MySQL using [KubeDB](https://kubedb.com/docs/latest/guides/mysql/) Scaling OpsRequest
  - Contact AppsCode team

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
    - If the database enters a `Critical` phase and the secondary pod goes `OFFLINE`, with pod logs showing the errors the `source's binary log is corrupted` and `Error_code: MY-013121`, it's likely that both the relay log and binary log are corrupted.
        - Here's the solution:
        ```bash
        # Step 1: Access the primary pod:
        $ kubectl exec -it -n <namespace> <primary-pod-name> -- bash 
        $ mysql -u<username> -p<password>
        $ SHOW MASTER STATUS;
        # From the query result, note down SOURCE_LOG_FILE and SOURCE_LOG_POS.
  
        # Step 2: Access the secondary (affected) pod:
        $ kubectl exec -it -n <namespace> <affected-pod-name> -- bash
        $ mysql -u<username> -p<password>
        $ STOP SLAVE;
        $ CHANGE REPLICATION SOURCE TO
          SOURCE_HOST = '<primary-pod-name>.<dbname>-pods.<namespace>.svc',
          SOURCE_PORT = 3306,
          SOURCE_USER = 'repl',
          SOURCE_PASSWORD = '$MYSQL_ROOT_PASSWORD',
          SOURCE_LOG_FILE = '<file>',
          SOURCE_LOG_POS = <position>;
        $ START SLAVE;
        ```
    - If any `MySQLOpsRequest` is ongoing on same database, Wait until it completes.
    - If some nodes of the MySQL group are not `Up`, Try restarting those nodes one at a time.
    - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If any `MySQLOpsRequest` is ongoing on same database, Wait until it completes.
  - Contact AppsCode team
- #### OpsRequestFailed
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
  - Check the `MySQL` CRs status
  - Contact AppsCode team
