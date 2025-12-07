package server

import (
	"bufio"
	"log"
	"math/rand"
	"net"
)

var random = rand.New(rand.NewSource(22))

type Client struct {
	id   uint
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		id:   uint(random.Uint64()),
		conn: conn,
	}
}

func (c *Client) handleRequest(channel chan Message, cleanup cleanUpFunc) {
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			cleanup(c.id)
			return
		}

		_, err = c.conn.Write([]byte("Message received.\n"))
		if err != nil {
			log.Println(err)
		}

		channel <- Message{content: message}
	}
}
