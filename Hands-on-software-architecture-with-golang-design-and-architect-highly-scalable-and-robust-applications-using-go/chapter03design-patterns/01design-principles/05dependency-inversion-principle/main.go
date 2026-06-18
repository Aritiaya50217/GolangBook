package main

type User struct{}

type UserRepository interface {
	FindByID(id uint) (*User, error)
}

type MySQLRepository struct{}

func (r *MySQLRepository) FindByID(id uint) (*User, error) {
	return &User{}, nil
}

type MongoRepository struct{}

func (r *MongoRepository) FindByID(id uint) (*User, error) {
	return &User{}, nil
}

// Usecase ไม่รู้ว่า Database จริงคืออะไร
type UserUsecase struct {
	repo UserRepository
}
