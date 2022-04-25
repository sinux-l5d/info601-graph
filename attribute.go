package main

type Attribute string

func (a *Attribute) GetValue() string {
	return string(*a)
}
