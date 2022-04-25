package main

import "fmt"

const (
	menu = `
		1. Add a relationship
		2. Add a relationship bidirectional
		3. Get a relationship
		4. Get all relationships
		5. Print the node
		6. Exit
		`
)

func PromptAction() {
}

func main() {

	graph := NewGraph()

	// MAISON

	m1 := NewInstance("m1")
	m2 := NewInstance("m2")
	m3 := NewInstance("m3")
	m4 := NewInstance("m4")
	m5 := NewInstance("m5")
	m6 := NewInstance("m6")

	m6.AddRelationshipBidirectional("voisin", m3)

	m3.AddRelationshipBidirectional("voisin", m2)
	m3.AddRelationshipBidirectional("voisin", m1)
	m3.AddRelationshipBidirectional("voisin", m4)

	m1.AddRelationshipBidirectional("voisin", m4)
	m1.AddRelationshipBidirectional("voisin", m5)

	m4.AddRelationshipBidirectional("voisin", m5)
	m4.AddAttribute("name", "m4")

	graph.AddNodes(m1, m2, m3, m4, m5, m6)

	// PERSONNE

	human := NewConcept("human")
	flo := NewInstanceType("flo", human)
	evan := NewInstanceType("evan", human)

	flo.AddRelationshipBidirectional("friend", evan)

	flo.AddRelationship("habite", m1)
	evan.AddRelationship("habite", m6)
	graph.AddNodes(flo, evan)

	graph.Print()

	// TEST MAISONS

	chemin, existe := m6.CheminOptiVers(m5, "voisin")

	fmt.Printf("M6 voisin de M5 ? %v\n", existe)

	fmt.Print("Chemin : ")
	for _, node := range chemin {
		fmt.Printf("%s ", node)
	}
	fmt.Println()

	// TEST PERSONNES

	_, ok := graph.estVoisin(evan, flo)

	fmt.Printf("Evan est voisin de Flo? %v\n", ok)
	// fmt.Print("Chemin : ")
	// for _, node := range chemin2 {
	// 	fmt.Printf("%s ", node)
	// }
	// fmt.Println()

	// distances := m6.CheminAccessible("voisin")
	// for k, v := range distances {
	// 	fmt.Printf("%s : %v\n", k, v)
	// }
}
