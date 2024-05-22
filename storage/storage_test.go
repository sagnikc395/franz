package storage

import (
	"fmt"
	"testing"
)

func TestStoragePush(t *testing.T) {
	s := NewMemoryStore()
	for i := 0; i < 1000; i++ {
		offset, err := s.Push([]byte("foo\nbar\nbaz"))
		if err != nil {
			t.Error(err)
		}
		data, err := s.Get(offset)
		if err != nil {
			t.Error(err)
		}

		fmt.Println(string(data))
	}
}
