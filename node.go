package main

import "fmt"

type INode interface {
	AddRelationship(name string, node INode)
	AddRelationshipBidirectional(name string, node INode)
	AUnCheminVers(node INode, relation string) bool
	CheminProfondeurVers(node INode, relation string) ([]INode, bool)
	CheminAccessible(relation string) map[string]Origine
	CheminOptiVers(node INode, relation string) ([]string, bool)
	Get(name string) []INode
	GetName() string
	Print()
	Attributes() map[string][]Attribute
	AttributesOf(name string) []Attribute
	AddAttribute(name string, value string)
	GetAttribute(name string) []Attribute
}

// a node in the sense of a knowledge graph. Shouldn't be used directly.
type Node struct {
	name          string
	relationships map[string][]INode
	attribute     map[string][]Attribute
}

// Constructor of node
func NewNode(name string) *Node {
	return &Node{
		name:          name,
		relationships: make(map[string][]INode),
		attribute:     make(map[string][]Attribute),
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

func (n *Node) AUnCheminVers(node INode, relation string) bool {
	pile := NewPile()
	pile.Push(n)
	visited := make(map[INode]bool)
	visited[n] = true
	for !pile.IsEmpty() {
		current := pile.Pop()
		if current == node {
			return true
		}
		for _, node := range current.Get(relation) {
			if !visited[node] {
				pile.Push(node)
				visited[node] = true
			}
		}
	}
	return false
}

func choisirProchain(n INode, relation string, visited map[string]bool) (INode, bool) {
	// fmt.Println("VOISINS de", n.GetName())
	for _, node := range n.Get(relation) {
		// fmt.Printf("%s : %v\n", node.GetName(), visited[node.GetName()])
		if !visited[node.GetName()] {
			return node, true
		}
	}
	return nil, false
}

// func printVisited(visited map[string]bool) {
// 	for node, b := range visited {
// 		fmt.Printf("%s : %v\n", node, b)
// 	}
// }

// retourne le chemin entre n et node, inclus
func (n *Node) CheminProfondeurVers(node INode, relation string) ([]INode, bool) {
	pile := NewPile()
	pile.Push(n)

	visited := make(map[string]bool)
	visited[n.GetName()] = true

	// var courant INode = n

	for !pile.IsEmpty() {
		courant, existe := choisirProchain(pile.Top(), relation, visited)
		//if existe {
		//	fmt.Println("courant : ", courant.GetName())
		//} else {
		//	fmt.Println("Pas de prochain")
		//}

		if courant != nil && courant.GetName() == node.GetName() {
			pile.Push(courant)
			return pile.AsArray(), true
		}

		if !existe {
			pile.Pop()
			// fmt.Printf("len pile : %d\n", pile.Len())
			// courant = pile.Top()
		} else {
			pile.Push(courant)
			visited[courant.GetName()] = true
		}
	}
	return nil, false
}

type Origine struct {
	nbIntermediaire int
	predecesseur    string
}

func (n *Node) CheminAccessible(relation string) map[string]Origine {
	distances := make(map[string]Origine)
	distances[n.GetName()] = Origine{0, ""}

	var courant []INode = []INode{n}
	var suivant []INode

	profondeur := 0

	for len(courant) > 0 {
		profondeur++
		for _, node := range courant {
			for _, nodeSuivant := range node.Get(relation) {
				if _, existe := distances[nodeSuivant.GetName()]; !existe {
					distances[nodeSuivant.GetName()] = Origine{profondeur, node.GetName()}
					suivant = append(suivant, nodeSuivant)
				}
			}
		}
		courant = suivant
		suivant = []INode{}
	}

	return distances
}

func (n *Node) CheminOptiVers(node INode, relation string) ([]string, bool) {
	distances := n.CheminAccessible(relation)
	//print distances
	result := []string{node.GetName()}

	// var courant Origine
	courant, existe := distances[node.GetName()]

	if !existe {
		return nil, false
	}

	for courant.predecesseur != "" {
		// fmt.Println(courant.predecesseur, "->", courant.nbIntermediaire)
		result = append([]string{courant.predecesseur}, result...)
		courant = distances[courant.predecesseur]
	}

	return result, true
}

func (n *Node) CardSousGraph(relation string) int {
	// dans mon idée, je récupère le dictionnaire des distances
	distances := n.CheminAccessible(relation)
	// on renvoie la taille du dictionnaire
	return len(distances) - 1

}

func (i *Node) Attributes() map[string][]Attribute {
	return i.attribute
}

func (i *Node) AttributesOf(name string) []Attribute {
	return i.attribute[name]
}

func (i *Node) AddAttribute(name string, value string) {
	i.attribute[name] = append(i.attribute[name], Attribute(value))
}

func (c *Node) GetAttribute(name string) []Attribute {
	return c.attribute[name]
}
