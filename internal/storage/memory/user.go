package memory

import (
	"fmt"
	"golang-blog-web/internal/models"
	"strings"
)

func (s *Storage) isUserUnique(username, email string, excludeID uint) bool {
	for id, user := range s.users {
		if id != excludeID && strings.EqualFold(strings.TrimSpace(user.Username), strings.TrimSpace(username)) && strings.EqualFold(strings.TrimSpace(user.Email), strings.TrimSpace(email)) {
			return false
		}
	}
	return true
}

func (s *Storage) CreateUser(newUser *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ok := s.isUserUnique(newUser.Username, newUser.Email, newUser.ID)
	if !ok {
		return fmt.Errorf("")
	}

	s.nextID++
	newUser.ID = s.nextID
	s.users[newUser.ID] = newUser

	return nil
}

func (s *Storage) GetUserByUsername(username string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("Пользователь не найден")
}
