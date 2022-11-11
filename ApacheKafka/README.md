<p align="center"><img src="../img/kafkalogo.png"></p>

* [Set up Kafka broker](#Set-up-Kafka-broker)  
   * [Data type mapping between Postgres and Cassandra](#Data-type-mapping-between-Postgres-and-Cassandra) 
* [References](#References)

# Set up Kafka broker
Here is current docker-compose.yml that is running in my Oracle VM CentOS7 based machine.  
```YAML
version: '3'
services:
  zookeeper:
    image: 'confluentinc/cp-zookeeper:7.0.1'
    container_name: zookeeper
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
  broker:
    image: 'confluentinc/cp-kafka:7.0.1'
    container_name: broker
    ports:
      - 9092:9092
      - 29093:29093
    depends_on:
      - zookeeper
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: 'zookeeper:2181'
      # Exposes 9092 for external connections to the broker
      # Use kafka:29092 for connections internal on the docker network
      # See https://rmoff.net/2018/08/02/kafka-listeners-explained/ for details
      KAFKA_LISTENERS: "INTERNAL://:29092,EXTERNAL://:29093,GLOBAL://:9092"
      KAFKA_ADVERTISED_LISTENERS: 'INTERNAL://broker:29092,EXTERNAL://localhost:29093,GLOBAL://192.168.56.103:9092'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT,GLOBAL:PLAINTEXT'
      KAFKA_INTER_BROKER_LISTENER_NAME: INTERNAL
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
  kafka-ui:
    container_name: kafka-ui
    image: provectuslabs/kafka-ui:latest
    ports:
      - 8080:8080
    depends_on:
      - zookeeper
      - broker
    environment:
      KAFKA_CLUSTERS_0_NAME: kc1
      KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: broker:29092
      KAFKA_CLUSTERS_0_METRICS_PORT: 9997
      #KAFKA_CLUSTERS_0_SCHEMAREGISTRY: http://schemaregistry0:8085
      # KAFKA_CLUSTERS_0_KAFKACONNECT_0_NAME: first
      KAFKA_CLUSTERS_0_KAFKACONNECT_0_ADDRESS: http://broker:8083
```

## Steps to go trough


___
## References
* [Apache Kafka® Quick Start](https://developer.confluent.io/quickstart/kafka-docker/)  
* [Kafka listeners explained](https://rmoff.net/2018/08/02/kafka-listeners-explained/)  
* [Baeldung Kafka docker connection](https://www.baeldung.com/kafka-docker-connection)  
* [Running Kafka in Docker machine](https://medium.com/@marcelo.hossomi/running-kafka-in-docker-machine-64d1501d6f0b)  
* [Example Kafka Compose files](https://github.com/provectus/kafka-ui/blob/master/documentation/compose/DOCKER_COMPOSE.md)  
