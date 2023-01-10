package queue

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func CreateNewNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}
