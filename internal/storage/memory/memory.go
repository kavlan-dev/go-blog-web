package memory

import (
	"go-blog-web/internal/model"
	"sync"
)

type storage struct {
	posts      map[uint]*model.Post
	users      map[uint]*model.User
	mu         sync.Mutex
	nextPostId uint
	nextUserId uint
}

func NewStorage() *storage {
	return &storage{
		posts: make(map[uint]*model.Post),
		users: make(map[uint]*model.User),
		mu:    sync.Mutex{},
	}
}
