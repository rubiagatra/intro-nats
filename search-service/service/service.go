package service

type Service struct{}

func NewService() *Service {
	return new(Service)
}