package factorypattern

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID            uuid.UUID
	From          time.Time
	To            time.Time
	HairDresserID uuid.UUID
}

func CreateBooking(from, to time.Time, hairDresserID uuid.UUID) (*Booking, error) {
	closingTime, _ := time.Parse(time.Kitchen, "17:00pm")
	if from.After(closingTime) {
		return nil, errors.New("no appointments after closing time")
	}
	return &Booking{
		HairDresserID: uuid.New(),
		ID:            uuid.New(),
		From:          from,
		To:            to,
	}, nil

}
