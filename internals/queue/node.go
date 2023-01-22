package queue

import (
	"bytes"
)

type NodeData []byte

type Node struct {
	data NodeData
	prev *Node
	next *Node
}

func (n *Node) Contains(data NodeData) bool {
	return bytes.Equal(n.data, data)
}

func NewNode(data NodeData) *Node {
	return &Node{
		data: data,
	}
}
