package service

import "go-blog-web/internal/model"

type postStorage interface {
	CreatePost(newPost *model.Post) error
	FindPosts() *[]model.Post
	FindPostById(id uint) (*model.Post, error)
	FindPostByTitle(title string) (*model.Post, error)
	UpdatePost(id uint, updatePost *model.Post) error
	DeletePost(id uint) error
}

func (s *service) CreatePost(newPost *model.Post) error {
	if err := newPost.Validate(); err != nil {
		return err
	}

	return s.storage.CreatePost(newPost)
}

func (s *service) AllPosts() *[]model.Post {
	return s.storage.FindPosts()
}

func (s *service) PostByID(id uint) (*model.Post, error) {
	return s.storage.FindPostById(id)
}

func (s *service) PostByTitle(title string) (*model.Post, error) {
	return s.storage.FindPostByTitle(title)
}

func (s *service) UpdatePost(id uint, updatePost *model.Post) error {
	if err := updatePost.Validate(); err != nil {
		return err
	}

	return s.storage.UpdatePost(id, updatePost)
}

func (s *service) DeletePost(id uint) error {
	return s.storage.DeletePost(id)
}
