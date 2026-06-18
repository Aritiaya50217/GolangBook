package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type UserBuilder struct {
	user User
}

func (b *UserBuilder) Name(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) Email(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

func main() {
	user := (&UserBuilder{}).Name("John").Email("test@gmail.com").Build()
	fmt.Println("user : ", user)
}
