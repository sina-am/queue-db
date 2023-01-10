package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sina-am/queue-db/internals/client"
	"github.com/sina-am/queue-db/internals/command"
)

var UsageMessage = `available commands are:
	enqueue <queueName> <data>
	dequeue <queueName>
	quit
`

func main() {
	var address = flag.String("h", "127.0.0.1:1212", "server address")
	reader := bufio.NewReader(os.Stdin)
	conn, err := client.Dial(*address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		fmt.Printf("[%s]:> ", *address)
		cmdData, err := reader.ReadBytes('\n')
		cmdData = bytes.Replace(cmdData, []byte("\n"), []byte(""), -1)

		if err != nil {
			log.Fatal(err)
		}
		if bytes.Equal(cmdData, []byte("quit")) {
			break
		}
		cmd, err := command.DecodeMessage(cmdData)
		if err != nil {
			fmt.Printf("invalid command %s\n%s", cmdData, UsageMessage)
			continue
		}

		response, err := conn.SendMessage(cmd)
		if err != nil {
			fmt.Print("Connection closed by peer\n")
			break
		}
		fmt.Printf("%s\n", response)
	}
}
