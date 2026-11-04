package db

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type BookOperationsService interface {
	ListByUser(userID string) []Book
}

type UserService struct {
	users map[string]User
	bs    BookOperationsService
}

func NewUserService(initial []User, bs BookOperationsService) *UserService {
	users := make(map[string]User)
	for _, u := range initial {
		users[u.ID] = u
	}

	return &UserService{users: users, bs: bs}
}

func (us *UserService) Get(id string) (*User, []Book, error) {
	u, ok := us.users[id]
	if !ok {
		return nil, nil, fmt.Errorf("no user found for id %s", id)
	}
	books := us.bs.ListByUser(id)

	return &u, books, nil
}

func (us *UserService) Exists(id string) error {
	_, ok := us.users[id]
	if !ok {
		return fmt.Errorf("no user found for id %s", id)
	}

	return nil
}

func (us *UserService) Upsert(u User) (User, error) {
	u.ID = uuid.NewString()
	us.users[u.ID] = u

	return u, nil
}
