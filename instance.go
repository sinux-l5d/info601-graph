package main

// A Instance is a node in the sense of a knowledge graph.
type Instance struct {
	Node
	attribute map[string][]string
}

func NewInstance(name string) *Instance {
	return &Instance{
		Node:      *NewNode(name),
		attribute: make(map[string][]string),
	}
}

func (i *Instance) Attributes() map[string][]string {
	return i.attribute
}

func (i *Instance) AttributesOf(name string) []string {
	return i.attribute[name]
}

func (i *Instance) AddAttribute(name string, value string) {
	i.attribute[name] = append(i.attribute[name], value)
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
}

func (c *Instance) CheminVers(node INode, relation string) ([]INode, bool) {
	return c.Node.CheminVers(node, relation)
}
