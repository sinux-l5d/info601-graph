package main

import "fmt"

func makeCity(g *Graph) {
	batiment := NewConcept("batiment")

	// Concept batiment public
	batimentPublic := NewConceptOf("Bâtiment public", batiment)

	// Concepts batiments privés
	batimentPrive := NewConceptOf("Bâtiment privé", batiment)
	maison := NewConceptOf("Maison", batimentPrive)
	appartement := NewConceptOf("Appartement", batimentPrive)
	chateau := NewConceptOf("Château", batimentPrive)
	hotel := NewConceptOf("Hôtel", batimentPrive)

	// Instances bâtiments publics
	salleFetes := NewInstanceOf("Salle des fêtes", batimentPublic)
	mairie := NewInstanceOf("Mairie", batimentPublic)
	gymnase := NewInstanceOf("Gymnase", batimentPublic)
	ecole := NewInstanceOf("École", batimentPublic)

	// Instances bâtiments privés

	mRouge := NewInstanceOf("Maison rouge", maison)
	mJaune := NewInstanceOf("Maison jaune", maison)
	mBleu := NewInstanceOf("Maison bleue", maison)
	mVerte := NewInstanceOf("Maison verte", maison)
	mRose := NewInstanceOf("Maison rose", maison)

	martinez := NewInstanceOf("Martinez", hotel)
	princes := NewInstanceOf("Hôtel des Princes", hotel)
	versaille := NewInstanceOf("Versaille", chateau)

	mJaune.AddRelationshipBidirectional("voisin", mRouge)
	mRose.AddRelationshipBidirectional("voisin", mJaune)
	mBleu.AddRelationshipBidirectional("voisin", mRose)
	mVerte.AddRelationshipBidirectional("voisin", mBleu)
	mVerte.AddRelationshipBidirectional("voisin", mRouge)
	mairie.AddRelationshipBidirectional("voisin", mVerte)
	ecole.AddRelationshipBidirectional("voisin", gymnase)

	// Concepts personnes
	personne := NewConcept("Personne")
	retraite := NewConceptOf("Retraité", personne)
	enfant := NewConceptOf("Enfant", personne)
	actif := NewConceptOf("Actif", personne)

	// Paul
	paul := NewInstanceOf("Paul", actif)
	paul.AddRelationship("habite", mJaune)
	paul.AddRelationship("travail à", gymnase)
	paul.AddRelationship("travail à", mairie)
	paul.AddRelationship("travail à", ecole)
	paul.AddAttribute("metier", "polyvalent")
	paul.AddAttribute("age", "30")

	// Jean
	jean := NewInstanceOf("Jean", actif)
	jean.AddRelationship("travail à", gymnase)
	jean.AddRelationship("habite", mRouge)
	jean.AddAttribute("metier", "Gardien")

	// Macron
	macron := NewInstanceOf("Macron", retraite)
	macron.AddRelationship("habite", versaille)
	macron.AddRelationship("travail à", mairie)
	macron.AddAttribute("metier", "Président")

	g.AddNodes(batiment, batimentPublic, batimentPrive, maison, appartement, chateau, hotel, salleFetes, mairie, gymnase, ecole, mRouge, mJaune, mBleu, mVerte, mRose, martinez, princes, versaille, personne, retraite, enfant, actif, paul, jean, macron)
}

func main() {
	graph := NewGraph()
	makeCity(graph)

	graph.Print()

	paulI, _ := graph.FindName("Paul")
	paul := paulI.(*Instance)
	jeanI, _ := graph.FindName("Jean")
	jean := jeanI.(*Instance)
	macronI, _ := graph.FindName("Macron")
	macron := macronI.(*Instance)

	_, sontVoisins := graph.estVoisin(paul, jean)

	fmt.Printf("%s est-il voisin de %s ? %v\n", paul.GetName(), jean.GetName(), sontVoisins)

	_, sontVoisins = graph.estVoisin(paul, macron)
	fmt.Printf("%s est-il voisin de %s ? %v\n", paul.GetName(), macron.GetName(), sontVoisins)

	habitatJean := jean.Get("habite")[0].(*Instance)
	habitatPaul := paul.Get("habite")[0].(*Instance)
	chemin, existe := habitatJean.CheminProfondeurVers(habitatPaul, "voisin")

	fmt.Printf("%s est-il voisin de %s ? %v\n", habitatPaul.GetName(), habitatJean.GetName(), existe)

	if existe {
		fmt.Printf("Le chemin est : %s\n", chemin)
	}
}

func main2() {

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
	flo := NewInstanceOf("flo", human)
	evan := NewInstanceOf("evan", human)

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
