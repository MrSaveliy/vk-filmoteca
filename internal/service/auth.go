package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/MrSaveliy/vk-filmoteca/internal/repository"
	"github.com/MrSaveliy/vk-filmoteca/models"
)

const (
	salt      = "adsfsdglls213lkdf21k2"
	signinKey = "asdf3232kj42^$^"
	tokenTTL  = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)

}

func (s *AuthService) CreateRole(role models.Role) (int, error) {
	return s.repo.CreateRole(role)

}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
