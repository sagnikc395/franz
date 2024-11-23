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
	_, err = s.Store.Fetch(-1)
	fmt.Println(err)
}
