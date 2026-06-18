package main

type User struct {
	Name string
}

type UserRepository interface {
	FindByID(id uint) (*User, error)
}

type MySQLRepository struct{}

func (r MySQLRepository) FindByID(id uint) (*User, error) {
	// query mysql
	return &User{}, nil
}

type MongoRepository struct{}

func (r MongoRepository) FindByID(id uint) (*User, error) {
	// query mongo
	return &User{}, nil
}

// เพิ่ม Repository ใหม่ได้โดยไม่ต้องแก้ Usecase
type Booking struct {
	Name string
}
type BookingRepository interface {
	FindBookingByID(id uint) (*Booking, error)
}

type UserUsecase struct {
	repo        UserRepository 
	bookingRepo BookingRepository // เพิ่มใหม่ สามารถเรียกใช้ได้เลย Ex booking.FindBookingByID(xxx)
}
