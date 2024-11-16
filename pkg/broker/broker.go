package broker

import (
	"sync"
	"time"

	"github.com/sagnikc395/franz/pkg/types"
)

type Broker struct {
	topics        map[string]*Topic
	subscribers   map[string][]*Subscribers
	mu            sync.RWMutex
	messageStrore MessageStore
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
