package main

import "fmt"

type INode interface {
	AddRelationship(name string, node INode)
	AddRelationshipBidirectional(name string, node INode)
	CheminVers(node INode, relation string) ([]INode, bool)
	Get(name string) []INode
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

func (n *Node) AddRelationshipBidirectional(name string, node INode) {
	n.AddRelationship(name, node)
	node.AddRelationship(name, n)
}

// Get relationships by name
func (n *Node) Get(name string) []INode {
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

func (n *Node) CheminVers(node INode, relation string) ([]INode, bool) {
	pile := NewPile()
	pile.Push(n)
	visited := make(map[INode]bool)
	visited[n] = true
	for !pile.IsEmpty() {
		current := pile.Pop()
		if current == node {
			return pile.AsArray(), true
		}
		for _, node := range current.Get(relation) {
			if !visited[node] {
				pile.Push(node)
				visited[node] = true
			}
		}
	}
}
