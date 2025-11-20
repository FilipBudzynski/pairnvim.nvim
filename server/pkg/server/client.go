package server

import (
	"bufio"
	"fmt"
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

func (c *Client) handleRequest() {
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			c.conn.Close()
			delete(G.clients, c.id)
			return
		}
		fmt.Printf("Message incoming: %s", string(message))

		// Send to all clients except self here
		_, err = c.conn.Write([]byte("Message received.\n"))
		if err != nil {
			log.Fatal(err)
		}

		for key, client := range G.clients {
			if key == c.id {
				continue
			}
			_, err = client.conn.Write([]byte(string(message)))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
