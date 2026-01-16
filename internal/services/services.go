package services

type StorageInterface interface {
	PostsStorage
	UsersStorage
}

type Service struct {
	storage StorageInterface
}

func New(storage StorageInterface) *Service {
	return &Service{storage: storage}
}
