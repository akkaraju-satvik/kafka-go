package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", "192.168.1.5:19092")
	if err != nil {
		log.Fatalln("failed to dial leader:", err)
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Fatalln(err, "failed to create topic")
	}

	topicConn, err := kafka.DialLeader(context.Background(), "tcp", "192.168.1.5:19092", topic, 0)
	if err != nil {
		log.Fatal(err)
	}

	topicConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	args := os.Args[1:]
	for _, arg := range args {
		_, err = topicConn.WriteMessages(
			kafka.Message{Value: []byte(arg)},
		)
		if err != nil {
			log.Fatalln("failed to write messages:", err)
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatalln("failed to close writer:", err)
	}

}
