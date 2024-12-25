package main

type LatLong struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocaltion() LatLong
	SetLocation(LatLong)
	CanFly() bool
	Spaek()
}

