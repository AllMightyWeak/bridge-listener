package main

import (
	"eth-event-listener/internal/listener"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	fmt.Println("Listener started")
	fmt.Println()
	listener.EthereumListener()

}
