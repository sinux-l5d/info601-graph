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

func NewConceptOf(name string, concept *Concept) *Concept {
	c := NewConcept(name)
	c.AddRelationship("rdf:subClassOf", concept)
	return c
}
