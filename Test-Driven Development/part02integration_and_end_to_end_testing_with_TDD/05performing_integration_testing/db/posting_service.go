package db

import "log"

type PostingService interface {
	NewOrder(b Book) error
}

type StubberPostingService struct{}

func NewPostingService() *StubberPostingService {
	return &StubberPostingService{}
}

func (sps *StubberPostingService) NewOrder(b Book) error {
	log.Printf("STUBBED POSTING SERVICE: book %s posted: %v", b.ID, b)
	return nil
}
