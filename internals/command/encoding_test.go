package command

import (
	"strings"
	"testing"
)

func TestInvalidOprationCommand(t *testing.T) {
	commandData := []byte("PUSH name {'name': 'sina'}")
	_, err := DecodeCommand(commandData)
	if err != ErrorInvalidOpration {
		t.Errorf("this is not a valid command.")
	}
}

func TestInsufficentData(t *testing.T) {
	commandData := []byte("enqueue name")
	_, err := DecodeCommand(commandData)
	if err != ErrorInvalidOpration {
		t.Errorf("this is not a valid command.")
	}
}

func TestRightCommand(t *testing.T) {
	op := "enqueue"
	name := "queue_name"
	data := "{\"user_id\": 321"
	commandData := []byte(strings.Join([]string{op, name, data}, " "))

	command, err := DecodeCommand(commandData)
	if err == ErrorInvalidOpration {
		t.Errorf("this is a valid command.")
	}
	if command.Opration != Enqueue || command.QueueName != name {
		t.Errorf("invalid")
	}

	if string(command.Data) != data {
		t.Error("invalid")
	}
}
