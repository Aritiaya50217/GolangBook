package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name string
	ID   string
	SSN  int
	Age  int
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("%s: %d,%s,%d",
		employee.Name,
		employee.Age, employee.ID,
		employee.SSN)
}

// SortByAge type
type SortByAge []Employee

func (sortIntf SortByAge) Len() int               { return len(sortIntf) }
func (sortIntf SortByAge) Swap(i, j int)          { sortIntf[i], sortIntf[j] = sortIntf[j], sortIntf[i] }
func (sortIntf SortByAge) Less(i int, j int) bool { return sortIntf[i].Age < sortIntf[j].Age }

// Mass and Miless Types
type Mass float64
type Miless float64

type Thing struct {
	name          string
	mass          Mass
	distance      Miless
	meltingpoint  int
	freezingpoint int
}

type ByFactor func(Thing1, Thing2 *Thing) bool

// Sort method
func (byFactor ByFactor) Sort(Things []Thing) {
	var sortedThings *ThingSorter
	sortedThings = &ThingSorter{
		Things:   Things,
		byFactor: byFactor,
	}
	sort.Sort(sortedThings)
}

// Len method
func (ThingSorter *ThingSorter) Len() int {
	return len(ThingSorter.Things)
}

// Swap method
func (ThingSorter *ThingSorter) Swap(i int, j int) {
	ThingSorter.Things[i], ThingSorter.Things[j] = ThingSorter.Things[j],
		ThingSorter.Things[i]
}

// Less method
func (ThingSorter *ThingSorter) Less(i int, j int) bool {
	return ThingSorter.byFactor(&ThingSorter.Things[i],
		&ThingSorter.Things[j])
}

// ThingSorter class
type ThingSorter struct {
	Things   []Thing
	byFactor func(Thing1 *Thing, Thing2 *Thing) bool
}

func main() {
	var employees = []Employee{
		{"Graham", "231", 235643, 31},
		{"John", "3434", 245643, 42},
		{"Michael", "8934", 32432, 17},
		{"Jenny", "24334", 32444, 26},
	}
	fmt.Println(employees)

	fmt.Println("===== sort by age =====")
	sort.Sort(SortByAge(employees))
	fmt.Println(employees)

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Age > employees[j].Age
	})
	fmt.Println(employees)

	fmt.Println("\n\n===== Things =====")
	var Things = []Thing{
		{"IronRod", 0.055, 0.4, 3000, -180},
		{"SteelChair", 0.815, 0.7, 4000, -209},
		{"CopperBowl", 1.0, 1.0, 60, -30},
		{"BrassPot", 0.107, 1.5, 10000, -456},
	}

	// by factor
	var name func(*Thing, *Thing) bool
	name = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.name < Thing2.name
	}

	var mass func(*Thing, *Thing) bool
	mass = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.mass < Thing2.mass
	}

	var distance func(*Thing, *Thing) bool
	distance = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.distance < Thing2.distance
	}

	var decreasingDistance func(*Thing, *Thing) bool
	decreasingDistance = func(p1, p2 *Thing) bool {
		return distance(p2, p1)
	}

	ByFactor(name).Sort(Things)
	fmt.Println("By name:", Things)
	ByFactor(mass).Sort(Things)
	fmt.Println("By mass:", Things)
	ByFactor(distance).Sort(Things)
	fmt.Println("By distance:", Things)
	ByFactor(decreasingDistance).Sort(Things)
	fmt.Println("By decreasing distance:", Things)
}
