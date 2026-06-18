package main

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
}

type HotelReservation interface {
	Reservation
	ChangeType()
}

type FlightReservation interface {
	Reservation
	AddExtraLuggageAllowance(peices int)
}

type HotelReservationImpl struct {
	reservationDate string
}

func (r HotelReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r HotelReservationImpl) CalculateCancellationFee() float64 {
	return 1.0
}

type FlightReservationImpl struct {
	reservationDate string
	luggageAllowed  int
}

func (r FlightReservationImpl) AddExtraLuggageAllowance(peices int) {
	r.luggageAllowed = peices
}

func (r FlightReservationImpl) CalculateCancellationFee() float64 {
	return 2.0
}

func (r FlightReservationImpl) GetReservationDate() string {
	return r.reservationDate
}
