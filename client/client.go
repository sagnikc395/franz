package client

import "bytes"

// client is  an instance of client connected to a set of franz servers
type Client struct {
	addrs []string
	buff  bytes.Buffer
}

// constructor for Client
func NewClient(addrs []string) *Client {
	return &Client{
		addrs: addrs,
	}
}

// SendMsg to send messages to the franz servers
func (s *Client) SendMsg(msgs []byte) error {
	_, err := s.buff.Write(msgs)
	return err
}

//Receive
