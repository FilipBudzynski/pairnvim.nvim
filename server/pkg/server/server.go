package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	config  Config
	clients map[uint]*Client
}

var G *Server

type Config struct {
	Host string
	Port string
}

func (c Config) toString() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func New(config Config) *Server {
	if G == nil {
		G = &Server{
			config:  config,
			clients: make(map[uint]*Client),
		}
	}

	return G
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.config.toString())
	if err != nil {
		log.Fatal("unable to create listener: %w", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		client := NewClient(conn)

		s.clients[client.id] = client
		go client.handleRequest()
	}
}
