package main

type User struct{}

type UserRepository interface {
	FindByID(id uint) (*User, error)
}

type MySQLRepo struct{}

func (r MySQLRepo) FindByID(id uint) (*User, error)

type MongoRepo struct{}

func (r MongoRepo) FindByID(id uint) (*User, error)

type MockRepo struct{}

func (r MockRepo) FindByID(id uint) (*User, error)

// Usecase ควรทำงานเหมือนเดิม ไม่ว่าจะ inject MySQLRepo , MongoRepo , MockRepo
type UserUsecase struct {
	repo UserRepository
}
