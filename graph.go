package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	// visited map[INode]bool
	nodes []INode
}

func NewGraph() *Graph {
	return &Graph{
		// visited: make(map[INode]bool),
		nodes: make([]INode, 0),
	}
}

func (g *Graph) AddNode(node INode) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) FindName(name string) []INode {
	var result []INode
	for _, node := range g.nodes {
		if node.GetName() == name {
			result = append(result, node)
		}
	}
	return result
}

func (g *Graph) ConceptOf(i *Instance) []Concept {
	var result []Concept
	concept := i.Get("rdf:type")
	if len(concept) > 1 || len(concept) == 0 {
		return result
	} else {
		result = append(result, *(concept[0].(*Concept)))
	}

	// while concept has rdf:subClassOf
	for {
		concept = concept[0].(*Concept).Get("rdf:subClassOf")
		if len(concept) > 1 || len(concept) == 0 {
			return result
		} else {
			result = append(result, *(concept[0].(*Concept)))
		}
	}
}

func (g *Graph) InstanceOf(c *Concept) []Instance {
	var result []Instance
	for _, node := range g.nodes {
		// Si c'est une instance
		if t, ok := node.(*Instance); ok {
			concepts := g.ConceptOf(t)
			// if c est dans concepts
			for _, concept := range concepts {
				if concept.GetName() == c.GetName() {
					result = append(result, *t)
				}
			}
		}
	}
	return result
}

func (g *Graph) StringInstanceOf(c *Concept) string {
	instances := g.InstanceOf(c)
	instancesS := ConvertInstanceToString(instances, func(item Instance) string {
		return item.GetName()
	})
	return fmt.Sprint(strings.Join(instancesS, " | ") + "\n")
}

func (g *Graph) StringConceptOf(i *Instance) string {
	concepts := g.ConceptOf(i)
	conceptsS := ConvertConceptToString(concepts, func(item Concept) string {
		return item.GetName()
	})
	// insert juliette at position 0²
	conceptsS = append([]string{i.GetName()}, conceptsS...)
	return fmt.Sprint(strings.Join(conceptsS, "->") + "\n")
}

func (g *Graph) estVoisin(a, b *Instance) ([]*Instance, bool) {
	habitationA := a.Get("habite")[0].(*Instance)
	habitationB := b.Get("habite")[0].(*Instance)

	chemin, existe := habitationA.CheminVers(habitationB, "habite")

	var result []*Instance
	for _, node := range chemin {
		result = append(result, node.(*Instance))
	}

	return result, existe
}

// print graph
func (g *Graph) Print() {
	for _, node := range g.nodes {
		node.Print()
	}
}