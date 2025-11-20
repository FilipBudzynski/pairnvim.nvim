package main

import "server/pkg/server"

func main() {
	server := server.New(server.Config{Host: "127.0.0.1", Port: "6666"})
	server.Run()
}
