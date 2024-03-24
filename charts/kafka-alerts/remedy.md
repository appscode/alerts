## TODOs on Kafka Critical Alerts

### Database Alerts

- #### KafkaUnderReplicatedPartitions:
  - Check the health of Kafka brokers and ensure they are reachable.
  - Adjust replication factors or increase the number of replicas for affected topics.
  - Monitor disk space and ensure it's not causing replication delays.
- #### KafkaAbnormalControllerState:
  - Restart the Kafka controller process.
  - Check for any issues with Zookeeper, as it is used for leader election.
- #### KafkaOfflinePartitions:
  - Restart Kafka brokers to bring offline partitions back online.
  - Investigate disk failures or network issues impacting partitions.
- #### KafkaUnderMinIsrPartitionCount:
  - Increase the min.insync.replicas configuration for affected topics.
  - Monitor network latency between brokers.
- #### KafkaOfflineLogDirectoryCount:
  - Check disk health and ensure log directories are accessible.
  - Increase disk space or clean up old logs to free up space.
- #### KafkaISRExpandRate and KafkaISRShrinkRate:
  - Monitor replication rates and adjust broker configurations if necessary.
  - Add more brokers to distribute the replication load.
- #### KafkaBrokerCount:
  - Ensure all expected Kafka brokers are up and running.
  - Investigate network or hardware issues impacting broker availability.
- #### KafkaNetworkProcessorIdlePercent and KafkaRequestHandlerIdlePercent:
  - Monitor CPU and network usage on Kafka brokers.
  - Adjust thread pools or configurations based on workload patterns.
- #### KafkaReplicaFetcherManagerMaxLag:
  - Monitor replica lag and adjust fetcher configurations.
  - Consider increasing network bandwidth or optimizing network settings.
- #### KafkaTopicCount:
  - Monitor topic creation and deletion activities.
  - Ensure proper resource allocation for managing topic metadata.
- #### KafkaPhaseCritical
  - Describe the `Kafka` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team
- #### KafkaDown:
  - Describe the `Kafka` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team

### KubeDB Provisioner

- #### KubeDBKafkaPhaseNotReady
  - Describe the `Kafka` CR, check the reason from conditions.
  - Check kafka pod logs
  - Check KubeDB provisioner logs
  - Contact AppsCode team
- #### KubeDBKafkaPhaseCritical
  - If any `KafkaOpsRequest` is progressing for the same database, wait until it completes.
  - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
  - Describe the pods that are not Ready and check their events and conditions. 
  - Contact AppsCode team if this persists for more than 30 minutes.

