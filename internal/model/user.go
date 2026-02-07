package model

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
	Role      string    `json:"role"` // user, admin
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	// Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Role string `json:"role"`
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
	if !isValidEmail(strings.TrimSpace(u.Email)) {
		return fmt.Errorf("Некорректный формат email")
	}

	return nil
}

func isValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	// Проверяем наличие @ и что она не первая и не последняя
	atIndex := strings.LastIndex(email, "@")
	if atIndex <= 0 || atIndex == len(email)-1 {
		return false
	}

	// Проверяем локальную часть (до @)
	localPart := email[:atIndex]
	if len(localPart) < 1 || len(localPart) > 64 {
		return false
	}

	// Проверяем доменную часть (после @)
	domainPart := email[atIndex+1:]
	if len(domainPart) < 1 {
		return false
	}

	// Проверяем наличие точки в домене
	if !strings.Contains(domainPart, ".") {
		return false
	}

	// Проверяем что домен не начинается или не заканчивается точкой
	if strings.HasPrefix(domainPart, ".") || strings.HasSuffix(domainPart, ".") {
		return false
	}

	// Проверяем что нет двух точек подряд
	if strings.Contains(domainPart, "..") {
		return false
	}

	return true
}

func (u *User) IsUserUnique(users []User) bool {
	for _, user := range users {
		if u.ID > 0 && user.ID == u.ID {
			continue
		}
		if strings.EqualFold(strings.TrimSpace(user.Username), strings.TrimSpace(u.Username)) {
			return false
		}
		if strings.EqualFold(strings.TrimSpace(user.Email), strings.TrimSpace(u.Email)) {
			return false
		}
	}
	return true
}
