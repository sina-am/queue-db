package main

import (
	"flag"
	"log"

	"github.com/sina-am/queue-db/internals/server"
)

func main() {
	var address = flag.String("address", "127.0.0.1:1212", "listening address.")
	flag.Parse()

	log.Printf("Server is binding on %s\n", *address)
	server.RunServer(*address)
}
