package tree

import "reflect"

// Node represents a tree node.
type Node struct {
	Value   interface{}
	Nodes   []*Node
	Virtual bool
}

// New creates a new tree / root node.
func New(value interface{}) *Node {
	return &Node{
		Value: value,
	}
}

// Add adds a leaf node
func (n *Node) Add(value interface{}) *Node {
	node := New(value)
	n.Nodes = append(n.Nodes, node)
	return node
}

// Find finds the child node with the target value.
// Nil if not found.
func (n *Node) Find(value interface{}) *Node {
	for _, node := range n.Nodes {
		if reflect.DeepEqual(node.Value, value) {
			return node
		}
	}
	return nil
}
