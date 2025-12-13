package main

import (
	"eth-event-listener/internal/listener"
	"fmt"
)

func main() {
	fmt.Println("Listener started")
	fmt.Println()
	listener.EthereumListener()

}
