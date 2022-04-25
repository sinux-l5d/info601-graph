package main

func main() {

	human := NewConcept("Human")

	woman := NewConcept("Woman")

	woman.AddRelationship("is-a", human)

	woman.Print()

	juliette := NewNode("Juliette")

	juliette.AddRelationship("is-a", woman)

	juliette.Print()

}
