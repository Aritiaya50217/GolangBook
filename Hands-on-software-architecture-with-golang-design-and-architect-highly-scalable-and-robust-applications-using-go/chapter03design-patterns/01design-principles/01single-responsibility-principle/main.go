package main

type UserService struct{}

func (s *UserService) CreateUser() error {
	return nil
}

type EmailService struct{}

func (s *EmailService) SendEmail() error {
	return nil
}

type ReportService struct{}

func (s *ReportService) GenerateReport() error {
	return nil
}
