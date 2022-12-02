// Commands are like:
// Opration QueueName [BinaryData]
// Opration should be "enqueue" or "dequeue".
// Example:
// enqueue orders {"user_id": 45421, "order_id": 634}
// dequque orders

package command

import (
	"bytes"
)

type opration string

const (
	Enqueue opration = "enqueue"
	Dequeue opration = "dequeue"
)

type Command struct {
	Opration  opration
	QueueName string
	Data      []byte
}

func EncodeCommand(cmd Command) []byte {
	return bytes.Join(
		[][]byte{[]byte(cmd.Opration), []byte(cmd.QueueName), cmd.Data},
		[]byte(" "))
}

func DecodeCommand(commandData []byte) (Command, error) {
	list := bytes.SplitN(commandData, []byte(" "), 3)
	if len(list) < 2 {
		return Command{}, ErrorInvalidOpration
	}

	command := Command{
		Opration:  opration(string(list[0])),
		QueueName: string(list[1]),
	}
	if command.Opration == Dequeue {
		return command, nil
	}
	if command.Opration == Enqueue && len(list) == 3 {
		command.Data = list[2]
		return command, nil
	}
	return command, ErrorInvalidOpration
}
