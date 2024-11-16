package main

import "github.com/sagnikc395/franz/pkg/storage"

func main() {
	//init message store
	store := storage.NewInMemoryStore()
	//create broker
	b := broker.NewBroker(store)
}
