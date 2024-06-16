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

func (mq *MessageQueue) Dequeue() (Message, bool) {
	// remove a message from the queue if it can be removed
	// else return a empty message and return false
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	if len(mq.queue) == 0 {
		return Message{}, false
	}
	msg := mq.queue[0]
	//slice off the item and rebase the queue
	mq.queue = mq.queue[1:]
	return msg, true
}
