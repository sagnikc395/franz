package main

import (
	"fmt"
	"log"

	"github.com/sagnikc395/franz/storage"
)

func main() {
	cfg := &Config{
		ListenAddr: ":3000",
		Store:      storage.NewMemoryStore(),
	}
	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	offset, _ := s.Store.Push([]byte("foobar"))
	data, err := s.Store.Fetch(offset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
