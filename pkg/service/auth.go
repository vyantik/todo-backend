package service

import (
	"crypto/sha1"
	"encoding/hex"
	"os"

	"github.com/vyantik/todo-backend"
	"github.com/vyantik/todo-backend/pkg/repository"
)

var salt = os.Getenv("SALT")

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	hash, err := s.generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = hash

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum([]byte(salt))), nil
}
