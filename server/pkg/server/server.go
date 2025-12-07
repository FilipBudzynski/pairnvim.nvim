package server

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	mu               *sync.RWMutex
	config           Config
	clients          map[uint]*Client
	broadcastChannel chan Message
}

type cleanUpFunc func(clientID uint)

type Message struct {
	content string
}

// I dont think its needed
// var singleInstance *Server
//
// func New(config Config) *Server {
// 	if singleInstance != nil {
// 		return singleInstance
// 	}
// 	singleInstance = &Server{
// 		config:  config,
// 		clients: make(map[uint]*Client),
// 	}
// 	return singleInstance
// }

func New(config Config) *Server {
	return &Server{
		mu:               &sync.RWMutex{},
		config:           config,
		clients:          make(map[uint]*Client),
		broadcastChannel: make(chan Message),
	}
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.config.toString())
	if err != nil {
		log.Fatal("unable to create listener: %w", err)
	}
	defer listener.Close()

	go s.broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		client := NewClient(conn)

		s.mu.Lock()
		s.clients[client.id] = client
		s.mu.Unlock()
		go client.handleRequest(s.broadcastChannel, s.clientCleanup)
	}
}

func (s *Server) broadcast() {
	for {
		msg := <-s.broadcastChannel
		s.mu.RLock()
		for _, client := range s.clients {
			_, err := client.conn.Write([]byte(msg.content))
			if err != nil {
				log.Println(err)
			}
		}
		s.mu.RUnlock()
	}
}

func (s *Server) clientCleanup(clientID uint) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	c := s.clients[clientID]
	_ = c.conn.Close()

	delete(s.clients, clientID)
}
