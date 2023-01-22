package main

import (
	"flag"
	"log"

	"github.com/sina-am/queue-db/internals/queue"
	"github.com/sina-am/queue-db/internals/server"
)

func main() {
	var (
		address   = flag.String("address", "127.0.0.1:1212", "listening address.")
		maxMemory = flag.Int("memory", 5000000, "Maximum memory usage for data (default ~5M)")
	)
	flag.Parse()

	queues := queue.NewMemoryQueueStorage()

	ops := server.TCPServerOps{
		Addr:         *address,
		MaxQueueSize: *maxMemory,
	}
	srv := server.NewTCPServer(ops, queues)
	log.Fatal(srv.Run())
}
