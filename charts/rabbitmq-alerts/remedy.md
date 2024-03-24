## TODOs on RabbitMQ Critical Alerts

### Database Alerts

- #### InsufficientEstablishedErlangDistributionLinks:
  - Check network connectivity between RabbitMQ nodes.
  - Ensure that the correct Erlang distribution ports are open between nodes.
  - Verify that nodes are using the correct hostnames or IP addresses to communicate.
  - Check permissions to ensure that nodes can communicate with each other.
- #### UnroutableMessages:
  - Review the exchange and queue bindings to ensure messages are routed correctly.
  - Check for any misconfigurations in routing rules.
  - Monitor for dead letter exchanges to handle unroutable messages appropriately.
- #### HighConnectionChurn:
  - Encourage clients to use long-lived connections instead of opening and closing connections frequently.
  - Review client connection handling logic to minimize unnecessary connections.
- #### LowDiskWatermarkPredicted:
  - Increase disk space allocation for RabbitMQ.
  - Monitor and limit queue backlogs to reduce disk usage.
  - Adjust message consumption rates to match available disk space.
  - Consider using auto-scaling mechanisms to increase disk size dynamically.
- #### FileDescriptorsNearLimit:
  - Reduce the number of connections, if possible.
  - Decrease the number of durable queues, if applicable.
  - Increase the file descriptor limit for RabbitMQ if necessary.
- #### TCPSocketsNearLimit:
  - Increase the number of channels to distribute connections more evenly across the cluster.
  - Implement connection pooling to reduce the number of TCP sockets used per connection.
  - Review client connection management to ensure connections are reused efficiently.

- #### RabbotMQDown
  - Describe the `RabbitMQ` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team

### KubeDB Provisioner

- #### KubeDBRabbitMQPhaseNotReady
  - Describe the `RabbitMQ` CR, check the reason from conditions.
  - Check rabbitmq pod logs
  - Check KubeDB provisioner logs
  - Contact AppsCode team
- #### KubeDBRabbitMQPhaseCritical
  - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
  - Describe the pods that are not Ready and check their events and conditions. 
  - Contact AppsCode team if this persists for more than 30 minutes.