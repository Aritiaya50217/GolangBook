package main

import "fmt"

type Event interface{}

type UserCreated struct {
	ID   string
	Name string
}

type UserRenamed struct {
	ID   string
	Name string
}

type State struct {
	Active bool
}

type User struct {
	ID   string
	Name string
}

func (u *User) Apply(event Event) {
	switch e := event.(type) {
	case UserCreated:
		u.ID = e.ID
		u.Name = e.Name
	case UserRenamed:
		u.Name = e.Name
	}
}

func main() {
	events := []Event{
		UserCreated{ID: "1", Name: "Alice"},
		UserRenamed{ID: "1", Name: "Alicia"},
	}

	var user User
	for _, e := range events {
		user.Apply(e)
	}

	fmt.Println(user)
}
