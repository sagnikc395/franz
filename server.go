package main

import "github.com/sagnikc395/franz/storage"

type Server struct {
	Store storage.Storer
	*Config
}

type Config struct {
	//default configuration for the server
	ListenAddr string
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
	}, nil
}

