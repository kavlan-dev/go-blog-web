package service

type storageInterface interface {
	postStorage
	userStorage
}

type service struct {
	storage storageInterface
}

func NewService(storage storageInterface) *service {
	return &service{storage: storage}
}
