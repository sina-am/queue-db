package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestEncodeMessage(t *testing.T) {
	msg := Message{
		Cmd:       Enqueue,
		QueueName: "queue",
		Data:      []byte("data"),
	}

	byteMsg := EncodeMessage(msg)
	assert.Equal(t, byteMsg, []byte("enqueue queue data"))

}

func BenchmarkEncodeMessage(b *testing.B) {
	msg := Message{
		Cmd:       Enqueue,
		QueueName: "queue",
		Data:      []byte("data"),
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			EncodeMessage(msg)
		}
	}
}

func BenchmarkDecodeMessage(b *testing.B) {
	byteMsg := []byte(`enqueue queue {"user_id": "1"}`)

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			DecodeMessage(byteMsg)
		}
	}
}
