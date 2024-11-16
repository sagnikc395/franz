package types

import "time"

type Message struct {
	ID        string
	Topic     string
	Payload   []byte
	Timestamp time.Time
	Sequence  uint64
}

type Subscription struct {
	ID       string
	Topic    string
	Consumer string
}

type Consumer struct {
	ID              string
	SubscriptionID  string
	LastProcessedID string
}
