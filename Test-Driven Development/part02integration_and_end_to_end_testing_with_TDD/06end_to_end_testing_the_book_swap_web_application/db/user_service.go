package db

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type BookOperationsService interface {
	ListByUser(userID string) ([]Book, error)
}

type UserService struct {
	db *gorm.DB
	bs BookOperationsService
}

func NewUserService(db *gorm.DB, bs BookOperationsService) *UserService {
	return &UserService{db: db, bs: bs}
}

func (us *UserService) Get(id string) (*User, []Book, error) {
	var u User
	if r := us.db.Where("id = ?", id).First(&u); r.Error != nil {
		return nil, nil, fmt.Errorf("no user found for id %s:%v", id, r.Error)
	}

	books, err := us.bs.ListByUser(id)
	if err != nil {
		return nil, nil, err
	}

	return &u, books, nil

}

func (us *UserService) Exists(id string) error {
	var u User
	if r := us.db.Where("id = ?", id).First(&u); r.Error != nil {
		return fmt.Errorf("no user found for id %s:%v", id, r.Error)
	}
	return nil
}

func (us *UserService) Upsert(u User) (User, error) {
	var eu User
	if r := us.db.Where("id = ?", u.ID).First(&eu); r.Error != nil {
		u.ID = uuid.NewString()
	}

	us.db.Save(&u)

	return u, nil
}
