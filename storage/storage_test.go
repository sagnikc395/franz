package storage_test

import (
	"fmt"
	"testing"

	"github.com/sagnikc395/franz/storage"
)

func TestStorage(t *testing.T) {
	s := storage.NewMemoryStore()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("foobarbaz_%d", i)
		offset, err := s.Push([]byte(key))
		if err != nil {
			t.Error(err)
		}

		data, err := s.Fetch(offset)
		if err != nil {
			t.Error(err)
		}

		fmt.Println(string(data))
	}
}
