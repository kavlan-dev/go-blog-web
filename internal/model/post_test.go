package model

import (
	"testing"
	"time"
)

func TestPostValidate(t *testing.T) {
	tests := []struct {
		name    string
		post    Post
		wantErr bool
	}{
		{
			name: "Valid post",
			post: Post{
				Title:   "Valid Title",
				Content: "Valid Content",
			},
			wantErr: false,
		},
		{
			name: "Empty title",
			post: Post{
				Title:   "",
				Content: "Valid Content",
			},
			wantErr: true,
		},
		{
			name: "Empty content",
			post: Post{
				Title:   "Valid Title",
				Content: "",
			},
			wantErr: true,
		},
		{
			name: "Title with only spaces",
			post: Post{
				Title:   "   ",
				Content: "Valid Content",
			},
			wantErr: true,
		},
		{
			name: "Content with only spaces",
			post: Post{
				Title:   "Valid Title",
				Content: "   ",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.post.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Post.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostIsTitleUnique(t *testing.T) {
	now := time.Now()
	posts := []Post{
		{
			ID:        1,
			Title:     "Post 1",
			Content:   "Content 1",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        2,
			Title:     "Post 2",
			Content:   "Content 2",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	tests := []struct {
		name       string
		post       Post
		wantUnique bool
	}{
		{
			name: "Unique title",
			post: Post{
				ID:      0,
				Title:   "New Post",
				Content: "New Content",
			},
			wantUnique: true,
		},
		{
			name: "Duplicate title",
			post: Post{
				ID:      0,
				Title:   "Post 1",
				Content: "New Content",
			},
			wantUnique: false,
		},
		{
			name: "Same post (update scenario)",
			post: Post{
				ID:      1,
				Title:   "Post 1",
				Content: "Updated Content",
			},
			wantUnique: true,
		},
		{
			name: "Case insensitive title",
			post: Post{
				ID:      0,
				Title:   "POST 1",
				Content: "New Content",
			},
			wantUnique: false,
		},
		{
			name: "Title with different case (update scenario)",
			post: Post{
				ID:      1,
				Title:   "post 1",
				Content: "Updated Content",
			},
			wantUnique: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.post.IsTitleUnique(posts); got != tt.wantUnique {
				t.Errorf("Post.IsTitleUnique() = %v, want %v", got, tt.wantUnique)
			}
		})
	}
}
