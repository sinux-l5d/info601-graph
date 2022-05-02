package main

import "fmt"

// A Instance is a node in the sense of a knowledge graph.
type Instance struct {
	Node
}

func NewInstance(name string) *Instance {
	return &Instance{
		Node: *NewNode(name),
	}
}

func NewInstanceOf(name string, concept *Concept) *Instance {
	i := NewInstance(name)
	i.AddRelationship("rdf:type", concept)
	return i
}

func (c *Instance) Print() {
	c.Node.Print()
	fmt.Println("Attributes:")
	for k, v := range c.Node.attribute {
		fmt.Println("\t", k, ":", v)
	}
}
