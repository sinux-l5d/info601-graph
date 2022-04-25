package main

import "fmt"

type INode interface {
	AddRelationship(name string, node INode)
	Find(name string) []INode
	GetName() string
	Print()
}

// a node in the sense of a knowledge graph. Shouldn't be used directly.
type Node struct {
	name          string
	relationships map[string][]INode
}

// Constructor of node
func NewNode(name string) *Node {
	return &Node{
		name:          name,
		relationships: make(map[string][]INode),
	}
}

func (n *Node) GetName() string {
	return n.name
}

// Add a relationship to the node
func (n *Node) AddRelationship(name string, node INode) {
	n.relationships[name] = append(n.relationships[name], node)
}

// Find relationships by name
func (n *Node) Find(name string) []INode {
	return n.relationships[name]
}

// Print the node and its relationships
func (n *Node) Print() {
	fmt.Println("[", n.name, "]")
	for relationName, nodes := range n.relationships {
		for _, node := range nodes {
			fmt.Println(" \u2514\u2500(", relationName, ")->[", node.GetName(), "]")
		}
	}
}
