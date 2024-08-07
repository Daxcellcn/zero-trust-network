package main

import (
	"log"
	"zero-trust-network/server" // This should be relative to the module path
)

func main() {
	log.Println("Starting Zero Trust Network server...")
	err := server.StartServer()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
