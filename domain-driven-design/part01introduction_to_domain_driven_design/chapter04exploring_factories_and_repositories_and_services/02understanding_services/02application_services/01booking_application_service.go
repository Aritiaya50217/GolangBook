package applicationservices

import (
	"context"
	factorypattern "domain-driven-design/part01introduction_to_domain_driven_design/04exploring_factories_and_repositories_and_services/01technical_requirements/01factory_pattern"
	repositorypattern "domain-driven-design/part01introduction_to_domain_driven_design/04exploring_factories_and_repositories_and_services/01technical_requirements/02repository_pattern"
	"errors"
	"fmt"
)

type accountKey = int

const accountCtxKey = accountKey(1)

type BookingDomainService interface {
	CreateBooking(ctx context.Context, booking factorypattern.Booking) error
}

type BookingAppService struct {
	bookingRepo          repositorypattern.BookingRepository
	bookingDomainService BookingDomainService
	emailService         EmailSender
}

func NewBookingAppService(bookingRepo repositorypattern.BookingRepository, bookingDomainService BookingDomainService, emailService EmailSender) *BookingAppService {
	return &BookingAppService{bookingRepo: bookingRepo, bookingDomainService: bookingDomainService, emailService: emailService}
}

func (b *BookingAppService) CreateBooking(ctx context.Context, booking factorypattern.Booking) error {
	u, ok := ctx.Value(accountCtxKey).(*Customer)
	if !ok {
		return errors.New("invalid customer")
	}

	if u.UserID != booking.UserID.String() {
		return errors.New("cannot create booking for other users")
	}

	if err := b.bookingDomainService.CreateBooking(ctx, booking); err != nil {
		return fmt.Errorf("could not create booking : %w", err)
	}

	if err := b.bookingRepo.SaveBooking(ctx, booking); err != nil {
		return fmt.Errorf("could not save booking : %w", err)
	}

	if err := b.emailService.SendEmail(ctx, "email", "booking", "success ..."); err != nil {
		return fmt.Errorf("cannot send email ...")
	}

	return nil
}
