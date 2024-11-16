package storage

import (
	"sync"

	"github.com/sagnikc395/franz/pkg/types"
)

type MessageStore interface {
	Store(msg types.Message) error
	Get(id string) (types.Message, error)
	GetRange(topic string, from, to uint64) ([]types.Message, error)
}

type InMemoryStore struct {
	messages map[string]types.Message
	mu       sync.RWMutex
}

// constructor for new memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		messages: make(map[string]types.Message),
	}
}

// implement store method
func (s *InMemoryStore) Store(msg types.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages[msg.ID] = msg
	return nil
}

// implement get method
func (s *InMemoryStore) Get(id string) (types.Message, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	msg, exists := s.messages[id]
	if !exists {
		return types.Message{}, MessageNotFoundErr
	}
	return msg, nil
}
