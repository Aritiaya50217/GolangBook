package main

import (
	"fmt"
	"sort"
)

type Commit struct {
	username string
	lang     string
	numlines int
}

type lessFunc func(p1, p2 *Commit) bool

type multiSorter struct {
	Commits      []Commit
	lessFunction []lessFunc
}

func (multiSorter *multiSorter) Sort(Commit []Commit) {
	multiSorter.Commits = Commit
	sort.Sort(multiSorter)
}

func OrderedBy(lessFunction ...lessFunc) *multiSorter {
	return &multiSorter{
		lessFunction: lessFunction,
	}
}

// Len method
func (multiSorter *multiSorter) Len() int {
	return len(multiSorter.Commits)
}

// Swap method
func (multiSorter *multiSorter) Swap(i int, j int) {
	multiSorter.Commits[i] = multiSorter.Commits[j]
	multiSorter.Commits[j] = multiSorter.Commits[i]
}

func (multiSorter *multiSorter) Less(i int, j int) bool {
	var p *Commit
	var q *Commit
	p = &multiSorter.Commits[i]
	q = &multiSorter.Commits[j]

	var k int
	for k = 0; k < len(multiSorter.lessFunction)-1; k++ {
		less := multiSorter.lessFunction[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return multiSorter.lessFunction[k](p, q)
}

func main() {
	var Commits = []Commit{
		{"james", "Javascript", 110},
		{"ritchie", "python", 250},
		{"fletcher", "Go", 300},
		{"ray", "Go", 400},
		{"john", "Go", 500},
		{"will", "Go", 600},
		{"dan", "C++", 500},
		{"sam", "Java", 650},
		{"hayvard", "Smalltalk", 180},
	}
	var user func(*Commit, *Commit) bool
	user = func(c1 *Commit, c2 *Commit) bool {
		return c1.username < c2.username
	}

	var language func(*Commit, *Commit) bool
	language = func(c1 *Commit, c2 *Commit) bool {
		return c1.lang < c2.lang
	}

	var increasingLines func(*Commit, *Commit) bool
	increasingLines = func(c1 *Commit, c2 *Commit) bool {
		return c1.numlines < c2.numlines
	}

	var decreasingLines func(*Commit, *Commit) bool
	decreasingLines = func(c1 *Commit, c2 *Commit) bool {
		return c1.numlines > c2.numlines // Note: > orders downwards.
	}
	
	OrderedBy(user).Sort(Commits)
	fmt.Println("By username:", Commits)
	OrderedBy(user, increasingLines).Sort(Commits)
	fmt.Println("By username,asc order", Commits)
	OrderedBy(user, decreasingLines).Sort(Commits)
	fmt.Println("By username,desc order", Commits)
	OrderedBy(language, increasingLines).Sort(Commits)
	fmt.Println("By lang,asc order", Commits)
	OrderedBy(language, decreasingLines, user).Sort(Commits)
	fmt.Println("By lang,desc order", Commits)
}
