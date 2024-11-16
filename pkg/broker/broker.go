package broker

import (
	"sync"
	"time"

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
