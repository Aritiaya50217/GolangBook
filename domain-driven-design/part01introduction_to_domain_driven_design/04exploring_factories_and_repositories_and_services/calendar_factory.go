package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	id            uuid.UUID
	from          time.Time
	to            time.Time
	hairDresserID uuid.UUID
}

func CreateBooking(from, to time.Time, hairDresserID uuid.UUID) (*Booking, error) {
	closingTime, _ := time.Parse(time.Kitchen, "17:00pm")
	if from.After(closingTime) {
		return nil, errors.New("no appointments after closing time")
	}
	return &Booking{
		hairDresserID: uuid.New(),
		id:            uuid.New(),
		from:          from,
		to:            to,
	}, nil

}
