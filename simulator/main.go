package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numServers := 100

	servers := make([]Server, numServers)

	for i := 0; i < numServers; i++ {
		servers[i] = Server{
			ID:     fmt.Sprintf("server-%05d", i+1),
			Status: "healthy",
		}
	}

	for {
		for _, server := range servers {
			event := generateEvent(server)

			fmt.Printf(
				"%s | CPU: %.2f | Memory: %.2f | Latency: %.2f | Status: %s\n",
				event.ServerID,
				event.CPU,
				event.Memory,
				event.Latency,
				event.Status,
			)
		}

		time.Sleep(time.Second)
	}
}