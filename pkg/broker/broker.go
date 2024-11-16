package broker

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sagnikc395/franz/pkg/storage"
	"github.com/sagnikc395/franz/pkg/types"
)

type Broker struct {
	topics        map[string]*Topic
	subscribers   map[string][]*Subscriber
	mu            sync.RWMutex
	messageStrore storage.MessageStore
}

type Topic struct {
	Name          string
	Messages      []types.Message
	LastSeqeunce  uint64
	RetentionTime time.Duration
	MaxSize       int
	mu            sync.RWMutex
}

type Subscriber struct {
	ID       string
	Topic    string
	Channel  chan types.Message
	Position uint64
}

// constructor for new broker
func NewBroker(messageStore storage.MessageStore) *Broker {
	return &Broker{
		topics:        make(map[string]*Topic),
		subscribers:   make(map[string][]*Subscriber),
		messageStrore: messageStore,
	}
}

// create a new topic
func (b *Broker) CreateTopic(name string, retention time.Duration, maxSize int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.topics[name]; exists {
		return TopicAlreadExistsErr
	}

	b.topics[name] = &Topic{
		Name:          name,
		RetentionTime: retention,
		MaxSize:       maxSize,
		Messages:      make([]types.Message, 0),
	}
	return nil
}

// publish
func (b *Broker) Publish(topic string, payload []byte) error {
	b.mu.RLock()
	t, exists := b.topics[topic]
	b.mu.RUnlock()

	if !exists {
		return TopicNotFoundErr
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	msg := types.Message{
		ID:        uuid.New().String(),
		Topic:     topic,
		Payload:   payload,
		Timestamp: time.Now(),
		Sequence:  t.LastSequence + 1,
	}

	t.LastSequence++
	t.Messages = append(t.Messages, msg)

	// Store message
	if err := b.messageStrore.Store(msg); err != nil {
		return err
	}

	// Broadcast to subscribers
	b.mu.RLock()
	subs := b.subscribers[topic]
	b.mu.RUnlock()

	for _, sub := range subs {
		select {
		case sub.Channel <- msg:
		default:
			// Channel full, implement backpressure handling here
		}
	}

	return nil
}

//sub on a given topic

func (b *Broker) Subscribe(topic string, bufferSize int) (*Subscriber, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.topics[topic]; !exists {
		return nil, TopicNotFoundErr
	}

	sub := &Subscriber{
		ID:      uuid.New().String(),
		Topic:   topic,
		Channel: make(chan types.Message, bufferSize),
	}

	b.subscribers[topic] = append(b.subscribers[topic], sub)
	return sub, nil
}

//unsub on a given topic

func (b *Broker) UnSubscribe(sub *Subscriber) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	subs := b.subscribers[sub.Topic]
	for i, s := range subs {
		if s.ID == sub.ID {
			close(s.Channel)
			b.subscribers[sub.Topic] = append(subs[:i], subs[i+1:]...)
			return nil
		}
	}

	return SubscriptionNotFoundErr
}
