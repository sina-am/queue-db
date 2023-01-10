package main

import (
	"flag"
	"log"

	"github.com/sina-am/queue-db/internals/queue"
	"github.com/sina-am/queue-db/internals/server"
)

func main() {
	var address = flag.String("address", "127.0.0.1:1212", "listening address.")
	flag.Parse()

	queues := queue.NewmemoryQueueList()
	srv := server.NewTCPServer(server.TCPServerOps{Addr: *address}, queues)
	log.Fatal(srv.Run())
}
