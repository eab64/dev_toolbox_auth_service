package usecase

import (
	"dev_toolbox_auth_service/internal/domain"
)

type AuthUseCase struct {
	userRepo domain.UserRepository
}

func NewAuthUseCase(userRepo domain.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
	}
}

func (a *AuthUseCase) Register(email, password, fullName string) error {
	user := &domain.User{
		Email:        email,
		PasswordHash: password, // потом добавим хеширование
		FullName:     fullName,
	}
	return a.userRepo.Create(user)
}
