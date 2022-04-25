package main

// A concept is a node in the sense of a knowledge graph.
type Concept struct {
	Node
}

func NewConcept(name string) *Concept {
	return &Concept{
		Node: *NewNode(name),
	}
}

func (c *Concept) GetName() string {
	return c.Node.GetName()
}

func (c *Concept) AddRelationship(name string, node INode) {
	c.Node.AddRelationship(name, node)
}

func (c *Concept) AddRelationshipBidirectional(name string, node INode) {
	c.Node.AddRelationshipBidirectional(name, node)
}

func (c *Concept) Get(name string) []INode {
	return c.Node.Get(name)
}

func (c *Concept) Print() {
	c.Node.Print()
}

func (c *Concept) CheminVers(node INode, relation string) ([]INode, bool) {
	return c.Node.CheminVers(node, relation)
}
