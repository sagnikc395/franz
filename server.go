package main

import (
	"net/http"

	"github.com/sagnikc395/franz/storage"
)

type Server struct {
	//	Store storage.Storer
	*Config
	topics map[string]storage.Storer
}

type Config struct {
	//default configuration for the server
	ListenAddr string
	Store      storage.Storer
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
		topics: make(map[string]storage.Storer),
	}, nil
}

func (s *Server) Start() {
	http.ListenAndServe(s.ListenAddr, nil)
}

func (s *Server) createTopic(name string)  {

}
