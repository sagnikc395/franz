package storage

import (
	"fmt"
	"log"
)

type Storer interface {
	// 2 functions every storage needs to apply on ,

	Push([]byte) (int, error)
	Get(int) ([]byte, error)
}

type MemoryStore struct {
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(data []byte) (int, error) {
	s.data = append(s.data, data)
	log.Printf("Added data to storage %v\n", s.data)
	return 0, nil
}

func (s *MemoryStore) Get(offset int) ([]byte, error) {
	if len(s.data) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}
	return s.data[offset], nil
}
