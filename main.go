package main

func main() {
	// Create a new node
	node := Node{
		Name: "A",
		Relationships: map[string][]*Node{
			"estAmi": {
				{Name: "B1"},
				{Name: "B2"},
			},
			"estEnnemi": {
				{Name: "C1"},
				{Name: "C2"},
			},
		},
	}

	// Print the node
	node.Print()

	node.Relationships["estAmi"][0].Print()

	// Add a new relationship
	// node.AddRelationship("D", &Node{Name: "D1"})

	// // Print the node
	// node.Print()

	// // Find the relationship
	// relationship := node.Find("D")

	// // Print the relationship
	// fmt.Println(relationship)
}
