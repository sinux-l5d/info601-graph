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

func (g *Graph) AddNodes(nodes ...INode) {
	g.nodes = append(g.nodes, nodes...)
}

func (g *Graph) FindName(name string) (INode, bool) {
	for _, node := range g.nodes {
		if node.GetName() == name {
			return node, true
		}
	}
	return nil, false
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

func (g *Graph) estVoisin(a, b *Instance) ([]INode, bool) {
	habitationA := a.Get("habite")[0].(*Instance)
	habitationB := b.Get("habite")[0].(*Instance)

	// fmt.Printf("%s habite %s\n", a.GetName(), habitationA.GetName())
	// fmt.Printf("%s habite %s\n", b.GetName(), habitationB.GetName())

	chemin, existe := habitationA.CheminProfondeurVers(habitationB, "voisin")

	if !existe {
		return nil, false
	}

	// // conversion
	// var result []*Instance
	// for _, node := range chemin {
	// 	result = append(result, node.(*Instance))
	// }

	return chemin, true
}

// print graph
func (g *Graph) Print() {
	for _, node := range g.nodes {
		node.Print()
		fmt.Println()
	}
}

// Fonction qui compte le nombre de quartiers
func (g *Graph) NbQuartier() int {
	//compteur := 0
	//trouver toutes les instances du concept batiment
	var instances []Instance
	var conceptBatiment INode
	var conceptTrouve bool
	var res []string
	var instanceEtVoisins []string
	conceptBatiment, conceptTrouve = g.FindName("batiment")
	if conceptTrouve {
		instances = g.InstanceOf(conceptBatiment.(*Concept))
	}
	//pour chaque instance de batiment
	for _, instance := range instances {
		//on prend tous ses voisins
		voisins := instance.CheminAccessible("voisin")
		// on ajoute l'instance et ses voisins à un tableau
		if !EstDans(instance.GetName(), instanceEtVoisins) {
			res = append(res, instance.GetName())
		}
		instanceEtVoisins = append(instanceEtVoisins, instance.GetName())
		for voisin := range voisins {
			instanceEtVoisins = append(instanceEtVoisins, voisin)
		}
	}
	return len(res)
}

// Fonction qui retourne les instances en fonction de leur attribut
func (g *Graph) GetInstancesPersonne(attribut string, relation string) []Instance {
	var instances []Instance
	var conceptPersonne INode
	var conceptTrouve bool
	conceptPersonne, conceptTrouve = g.FindName("Personne")
	if conceptTrouve {
		instances = g.InstanceOf(conceptPersonne.(*Concept))
	}
	var res []Instance
	for _, instance := range instances {
		if len(instance.GetAttribute(relation)) > 0 {
			if instance.GetAttribute(relation)[0].GetValue() == attribut {
				res = append(res, instance)
			}
		} else {
			// On récupère les concepts de l'instance
			concepts := g.ConceptOf(&instance)
			// On recherche l'attribut dans les concepts
			for _, concept := range concepts {
				if len(concept.GetAttribute(relation)) > 0 {
					if concept.GetAttribute(relation)[0].GetValue() == attribut {
						res = append(res, instance)
						break
					}
				}
			}
		}
	}
	return res
}
