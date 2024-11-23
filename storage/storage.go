package storage

import (
	"fmt"
	"sync"
)

type Storer interface {
	// store arbitary data and it will return an error
	Push([]byte) (int, error)
	// fetch data from an certain offset; return error if out of bounds
	Fetch(int) ([]byte, error)
}

type MemoryStore struct {
	mu   sync.RWMutex
	data [][]byte
}

// constructor
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

// push method
func (s *MemoryStore) Push(b []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

// fetch method
func (s *MemoryStore) Fetch(offset int) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset < 0 {
		return nil, fmt.Errorf("offset cannot be smaller than 0")
	}
	if len(s.data) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}

	return s.data[offset], nil
}
