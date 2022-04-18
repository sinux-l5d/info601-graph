package main

func main() {

	node := NewNode("Paul")

	node.AddRelationship("is-a", NewNode("Humain"))

	node.Print()

}
