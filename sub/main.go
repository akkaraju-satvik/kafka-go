package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	for {
		n, err := conn.ReadMessage(1e6)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(n.Value))
	}

}
