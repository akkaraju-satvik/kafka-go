package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {

	envFlag := flag.String("env", ".env", ".env file to be used")
	flag.Parse()

	godotenv.Load(*envFlag)

	topic := "my-topic"
	partition := 0

	kafkaEndpoint := os.Getenv("KAFKA_ENDPOINT")
	if kafkaEndpoint == "" {
		log.Fatalf("invalid kafka endpoint")
	}
	log.Println("Connecting to " + kafkaEndpoint)

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaEndpoint, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	for {
		n, err := conn.ReadMessage(1e6)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(n.Key), "-", string(n.Value), "-", n.Time)
	}

}
