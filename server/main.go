package main

import "github.com/FilipBudzynski/pairnvim.nvim.git/pkg/server"

func main() {
	config := server.Config{Host: "127.0.0.1", Port: "6666"}
	server := server.New(config)
	server.Run()
}
