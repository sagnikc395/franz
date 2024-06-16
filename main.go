package main

import "sync"

type Message struct {
	Data string
}

type MessageQueue struct {
	queue []Message
	mutex sync.Mutex
}

func (mq *MessageQueue) Enqueue(msg Message) {
	//add a item to the message queue
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	mq.queue = append(mq.queue, msg)
}
