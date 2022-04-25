package main

func ConvertConceptToString(vs []Concept, f func(Concept) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func ConvertInstanceToString(vs []Instance, f func(Instance) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

type Pile struct {
	pile []INode
}

func NewPile() *Pile {
	return &Pile{}
}

func (p *Pile) Push(n INode) {
	p.pile = append(p.pile, n)
}

func (p *Pile) Top() INode {
	return p.pile[len(p.pile)-1]
}

func (p *Pile) Pop() (top INode) {
	top = p.Top()
	p.pile = p.pile[:len(p.pile)-1]
	return
}

func (p *Pile) IsEmpty() bool {
	return len(p.pile) == 0
}

func (p *Pile) AsArray() []INode {
	return p.pile
}
