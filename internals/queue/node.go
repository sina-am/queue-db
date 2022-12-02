package queue

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func CreateNewNode(data interface{}) *Node {
	n := new(Node)
	n.data = data
	return n
}
