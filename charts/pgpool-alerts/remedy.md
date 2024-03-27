## TODOs on Pgpool Critical Alerts

### Database Alerts

- #### PgpoolDown
    - Describe the `Pgpool` CR, check the reason from conditions and try restarting the pods
    - Check pgpool pod logs
    - Contact AppsCode team
- #### PgpoolExporterLastScrapeError
    - Check logs for exporter containers  
    - Check if Pgpool is pingable
- #### PgpoolPostgresHealthCheckFailure
    - Check logs for backend postgresql database pods
    - Check if postgresql is pingable
- #### PgpoolTooManyConnections
    - Increase `max_pool` value in `initConfig` field using [PgpoolOpsRequest](https://kubedb.com/docs/latest/guides/pgpool/reconfigure/overview/)
- #### PgpoolBackendPanicMessageCount | PgpoolBackendFatalMessageCount | PgpoolBackendErrorMessageCount
    - Check logs for pgpool pods
    - Check logs for backend postgresql database pods
    - Check provisioner logs

### KubeDB Provisioner

- #### KubeDBPgpoolPhaseNotReady
    - Describe the `Pgpool` CR, check the reason from conditions.
    - Check pgpool pod logs
    - Check KubeDB provisioner logs
    - Contact AppsCode team
- #### KubeDBPgpoolPhaseCritical
    - If any `PgpoolDBOpsRequest` is progressing for the same pgpool, wait until it completes.
    - If some pods of the pgpool are not `Ready`, Try restarting those pods one at a time.
    - Describe the pods that are not Ready and check their events and conditions.
    - Contact AppsCode team if this persists for more than 30 minutes.
