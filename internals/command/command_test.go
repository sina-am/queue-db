package command

import (
	"strings"
	"testing"
)

func TestInvalidCmdMessage(t *testing.T) {
	commandData := []byte("PUSH name {'name': 'sina'}")
	_, err := DecodeMessage(commandData)
	if err != ErrInvalidMessage {
		t.Errorf("this is not a valid command.")
	}
}

func TestInsufficentData(t *testing.T) {
	commandData := []byte("enqueue name")
	_, err := DecodeMessage(commandData)
	if err != ErrInvalidMessage {
		t.Errorf("this is not a valid command.")
	}
}

func TestRightMessage(t *testing.T) {
	op := "enqueue"
	name := "queue_name"
	data := "{\"user_id\": 321"
	commandData := []byte(strings.Join([]string{op, name, data}, " "))

	command, err := DecodeMessage(commandData)
	if err == ErrInvalidMessage {
		t.Errorf("this is a valid command.")
	}
	if command.Cmd != Enqueue || command.QueueName != name {
		t.Errorf("invalid")
	}

	if string(command.Data) != data {
		t.Error("invalid")
	}
}
