package service

import (
	"github.com/DedeBKc/golang_api/dto"
	"github.com/DedeBKc/golang_api/entity"
	"github.com/DedeBKc/golang_api/repository"
	"github.com/jinzhu/gorm"
)

// AuthService apa yang dapa dilakukan
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
}

type authService struct {
	userRepository repository.UserRepository
}
