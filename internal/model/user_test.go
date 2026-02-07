package model

import (
	"testing"
	"time"
)

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "Valid user",
			user: User{
				Username: "validuser",
				Password: "validpassword",
				Email:    "valid@example.com",
			},
			wantErr: false,
		},
		{
			name: "Empty username",
			user: User{
				Username: "",
				Password: "validpassword",
				Email:    "valid@example.com",
			},
			wantErr: true,
		},
		{
			name: "Empty password",
			user: User{
				Username: "validuser",
				Password: "",
				Email:    "valid@example.com",
			},
			wantErr: true,
		},
		{
			name: "Empty email",
			user: User{
				Username: "validuser",
				Password: "validpassword",
				Email:    "",
			},
			wantErr: true,
		},
		{
			name: "Invalid email - no @",
			user: User{
				Username: "validuser",
				Password: "validpassword",
				Email:    "invalid.email.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserIsUserUnique(t *testing.T) {
	now := time.Now()
	users := []User{
		{
			ID:        1,
			Username:  "user1",
			Email:     "user1@example.com",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        2,
			Username:  "user2",
			Email:     "user2@example.com",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	tests := []struct {
		name       string
		user       User
		wantUnique bool
	}{
		{
			name: "Unique user",
			user: User{
				ID:       0,
				Username: "newuser",
				Email:    "new@example.com",
			},
			wantUnique: true,
		},
		{
			name: "Duplicate username",
			user: User{
				ID:       0,
				Username: "user1",
				Email:    "new@example.com",
			},
			wantUnique: false,
		},
		{
			name: "Duplicate email",
			user: User{
				ID:       0,
				Username: "newuser",
				Email:    "user1@example.com",
			},
			wantUnique: false,
		},
		{
			name: "Same user (update scenario)",
			user: User{
				ID:       1,
				Username: "user1",
				Email:    "user1@example.com",
			},
			wantUnique: true,
		},
		{
			name: "Case insensitive username",
			user: User{
				ID:       0,
				Username: "USER1",
				Email:    "new@example.com",
			},
			wantUnique: false,
		},
		{
			name: "Case insensitive email",
			user: User{
				ID:       0,
				Username: "newuser",
				Email:    "USER1@EXAMPLE.COM",
			},
			wantUnique: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.IsUserUnique(users); got != tt.wantUnique {
				t.Errorf("User.IsUserUnique() = %v, want %v", got, tt.wantUnique)
			}
		})
	}
}
