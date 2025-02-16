package main

import "fmt"

type LatLong struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocaltion() LatLong
	SetLocation(LatLong)
	CanFly() bool
	GetName() string
	Speak() string
}

type Animals struct {
	Lion   Lion
	Pigeon Pigeon
}

// The Lion Family
type Lion struct {
	name       string
	maneLength int
	location   LatLong
}

func (lion *Lion) GetLocaltion() LatLong {
	return lion.location
}

func (lion *Lion) SetLocation(loc LatLong) {
	lion.location = loc
}

func (lion *Lion) CanFly() bool {
	return false
}

func (lion *Lion) Spaek() string {
	return "roar"
}

func (lion *Lion) GetManeLength() int {
	return lion.maneLength
}

func (lion *Lion) GetName() string {
	return lion.name
}

// The Pigeon Family
type Pigeon struct {
	name     string
	location LatLong
}

func (p *Pigeon) GetLocaltion() LatLong {
	return p.location
}

func (p *Pigeon) SetLocation(loc LatLong) {
	p.location = loc
}

func (p *Pigeon) CanFly() bool {
	return false
}

func (p *Pigeon) Spaek() string {
	return "hoot"
}

func (p *Pigeon) GetName() string {
	return p.name
}

// The symphony
func makeThemSing(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetName() + " says " + animal.Speak())
	}
}

func main() {
	var aAnimal Animals
	aAnimal = Animals{
		Lion: Lion{
			name:       "Leo",
			maneLength: 10,
			location:   LatLong{10.40, 11.5},
		},
		Pigeon: Pigeon{
			name:     "Pigeon",
			location: LatLong{10.40, 11.5},
		},
	}
	fmt.Println(aAnimal)

}
