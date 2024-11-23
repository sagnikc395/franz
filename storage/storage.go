package storage

import "fmt"

type Storer interface {
	// store arbitary data and it will return an error
	Push([]byte) (int, error)
	// fetch data from an certain offset; return error if out of bounds
	Fetch(uint) ([]byte, error)
}

type MemoryStore struct {
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
	s.data = append(s.data, b)
	return len(s.data), nil
}

// fetch method
func (s *MemoryStore) Fetch(offset uint) ([]byte, error) {
	if uint(len(s.data)) < offset || offset == 0 {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}

	return s.data[offset-1], nil
}
