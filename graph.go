package main

import "fmt"

type Node struct {
	Name          string
	Relationships map[string][]*Node
}

func (n *Node) AddRelationship(name string, node *Node) {
	n.Relationships[name] = append(n.Relationships[name], node)
}

func (n *Node) Find(name string) []*Node {
	return n.Relationships[name]
}

func (n *Node) Print() {
	fmt.Println("[", n.Name, "]")
	for relationName, nodes := range n.Relationships {
		for _, node := range nodes {
			fmt.Println(" \u2514\u2500(", relationName, ")->[", node.Name, "]")
		}
	}
}
