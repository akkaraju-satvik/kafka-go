package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {

	envFlag := flag.String("env", ".env", ".env file to be used")
	flag.Parse()

	godotenv.Load(*envFlag)

	topic := "my-topic"
	kafkaEndpoint := os.Getenv("KAFKA_ENDPOINT")
	if kafkaEndpoint == "" {
		log.Fatalf("invalid kafka endpoint")
	}
	conn, err := kafka.Dial("tcp", kafkaEndpoint)
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

	topicConn, err := kafka.DialLeader(context.Background(), "tcp", kafkaEndpoint, topic, 0)
	if err != nil {
		log.Fatal(err)
	}

	topicConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	args := flag.Args()
	for _, arg := range args {
		_, err = topicConn.WriteMessages(
			kafka.Message{
				Key:   []byte(arg),
				Value: []byte(arg),
			},
		)
		if err != nil {
			log.Fatalln("failed to write messages:", err)
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatalln("failed to close writer:", err)
	}

}
