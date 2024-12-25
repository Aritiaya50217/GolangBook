package main

import "fmt"

type Pigeon struct {
	Name          string
	featherLength int
}

func (p *Pigeon) GetFeatherLergth() int {
	return p.featherLength
}

func (p *Pigeon) SetFeatherLegth(length int) {
	p.featherLength = length
}

func main() {
	p := Pigeon{
		Name:          "Tweety",
		featherLength: 10,
	}
	p.SetFeatherLegth(10)
	fmt.Println(p.Name)
	fmt.Println(p.featherLength)
}
