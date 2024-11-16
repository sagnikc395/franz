package consumer

import (
	"context"
	"sync"

	"github.com/sagnikc395/franz/pkg/broker"
	"github.com/sagnikc395/franz/pkg/types"
)

type Consumer struct {
	ID            string
	Topic         string
	Group         string
	subscription  *broker.Subscriber
	processor     MessageProcessor
	commitManager *CommitManager
	mu            sync.RWMutex
}

type MessageProcessor interface {
	Process(msg types.Message) error
}

type CommitManager struct {
	lastCommitted uint64
	mu            sync.Mutex
}

//constructor for new consumer

func NewConsumer(id, topic, group string, processor MessageProcessor) *Consumer {
	return &Consumer{
		ID:            id,
		Topic:         topic,
		Group:         group,
		processor:     processor,
		commitManager: &CommitManager{},
	}
}

// start the consumer
func (c *Consumer) Start(ctx context.Context) error {
	for {
		select {
		case msg := <-c.subscription.Channel:
			if err := c.processMessage(msg); err != nil {
				continue
			}
		case <-ctx.Done():
			return nil
		}
	}
}

// process the message
func (c *Consumer) processMessage(msg types.Message) error {
	//ensure the message ordering
	c.mu.Lock()
	defer c.mu.Unlock()

	if msg.Sequence <= c.commitManager.lastCommitted {
		return nil // as it is already processed
	}

	if err := c.processor.Process(msg); err != nil {
		return err
	}

	c.commitManager.commit(msg.Sequence)
	return nil
}
