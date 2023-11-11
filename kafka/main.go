package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

var (
	topic = "HVSE"
)

func main() {
	go func() {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "foo",
			"auto.offset.reset": "smallest"})
		if err != nil {
			log.Fatal(err)
		}
		err = consumer.Subscribe(topic, nil)
		if err != nil {
			log.Fatal(err)
		}
		for {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				// application-specific processing
				fmt.Printf("comsumed message from the queue: %s\n", string(e.Value))
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				break
			}
		}
		consumer.Close()
	}()

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"client.id":         "something",
			"acks":              "all",
		})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	deliveryChan := make(chan kafka.Event, 10000)
	for {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("FOO"),
		},
			deliveryChan,
		)
		if err != nil {
			log.Fatal(err)
		}

		<-deliveryChan
		time.Sleep(time.Second * 2)
	}
}
