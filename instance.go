package main

import "fmt"

// A Instance is a node in the sense of a knowledge graph.
type Instance struct {
	Node
	attribute map[string][]Attribute
}

func NewInstance(name string) *Instance {
	return &Instance{
		Node:      *NewNode(name),
		attribute: make(map[string][]Attribute),
	}
}

func NewInstanceType(name string, concept *Concept) *Instance {
	i := NewInstance(name)
	i.AddRelationship("rdf:type", concept)
	return i
}

func (i *Instance) Attributes() map[string][]Attribute {
	return i.attribute
}

func (i *Instance) AttributesOf(name string) []Attribute {
	return i.attribute[name]
}

func (i *Instance) AddAttribute(name string, value string) {
	i.attribute[name] = append(i.attribute[name], Attribute(value))
}

func (c *Instance) GetName() string {
	return c.Node.GetName()
}

func (c *Instance) AddRelationship(name string, node INode) {
	c.Node.AddRelationship(name, node)
}

func (c *Instance) AddRelationshipBidirectional(name string, node INode) {
	c.Node.AddRelationshipBidirectional(name, node)
}

func (c *Instance) Get(name string) []INode {
	return c.Node.Get(name)
}

func (c *Instance) Print() {
	c.Node.Print()
	fmt.Println("Attributes:")
	for k, v := range c.attribute {
		fmt.Println("\t", k, ":", v)
	}
}

func (c *Instance) AUnCheminVers(node INode, relation string) bool {
	return c.Node.AUnCheminVers(node, relation)
}

func (c *Instance) CheminProfondeurVers(node INode, relation string) ([]INode, bool) {
	return c.Node.CheminProfondeurVers(node, relation)
}

func (c *Instance) CheminAccessible(relation string) map[string]Origine {
	return c.Node.CheminAccessible(relation)
}

func (c *Instance) CheminOptiVers(node INode, relation string) ([]string, bool) {
	return c.Node.CheminOptiVers(node, relation)
}
