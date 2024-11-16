package main

import (
	"context"
	"log"
	"time"

	"github.com/sagnikc395/franz/pkg/broker"
	"github.com/sagnikc395/franz/pkg/consumer"
	"github.com/sagnikc395/franz/pkg/storage"
	"github.com/sagnikc395/franz/pkg/types"
)

func main() {
	//init message store
	store := storage.NewInMemoryStore()
	//create broker
	b := broker.NewBroker(store)
	// Create topic
	if err := b.CreateTopic("events", 24*time.Hour, 1000000); err != nil {
		log.Fatal(err)
	}

	// Create subscriber
	_, err := b.Subscribe("events", 100)
	if err != nil {
		log.Fatal(err)
	}

	// Create consumer
	processor := &MyMessageProcessor{}
	c := consumer.NewConsumer("consumer-1", "events", "group-1", processor)

	// Start consumer
	ctx := context.Background()
	go func() {
		if err := c.Start(ctx); err != nil {
			log.Printf("Consumer error: %v", err)
		}
	}()

	// Publish messages
	for i := 0; i < 10; i++ {
		if err := b.Publish("events", []byte("test message")); err != nil {
			log.Printf("Publish error: %v", err)
		}
	}

	// Wait for messages to be processed
	time.Sleep(time.Second)
}

type MyMessageProcessor struct{}

func (p *MyMessageProcessor) Process(msg types.Message) error {
	log.Printf("Processing message: %s", string(msg.Payload))
	return nil
}
