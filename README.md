# Kafka Golang Example

## Description

Kafka example with `segmentio/kafka-go` and `apache/kafka:3.7.0` docker image.

With Apache Kafka 3.7.0, Apache has released the official docker image for Kafka. This example uses the official docker image to run a Kafka broker without any additional configuration. `kafka-go` is a native Go client for Kafka that provides both low-level and high-level API support.

The docker-compose file is from the [example](https://github.com/apache/kafka/blob/trunk/docker/examples/jvm/single-node/plaintext/docker-compose.yml) provided in the official [Kafka repository](https://github.com/apache/kafka).
