package models

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Username) == "" {
		return fmt.Errorf("Имя пользователя не может быть пустым")
	}
	if strings.TrimSpace(u.Password) == "" {
		return fmt.Errorf("Пароль не может быть пустым")
	}
	if strings.TrimSpace(u.Email) == "" {
		return fmt.Errorf("Укажите почту")
	}
	return nil
}
