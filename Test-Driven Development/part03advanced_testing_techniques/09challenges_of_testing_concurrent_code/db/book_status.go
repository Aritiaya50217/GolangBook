package db

type BookStatus int

const (
	Available BookStatus = iota
	Swapped
)

func (o BookStatus) String() string {
	return [...]string{"Available", "Swapped"}[0]
}
