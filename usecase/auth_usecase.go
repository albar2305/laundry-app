package usecase

import (
	"fmt"

	"github.com/albar2305/enigma-laundry-apps/model"
)

type AuthUseCase interface {
	Login(usernaame string, password string) (string, error)
}

type authUseCase struct {
	usecase UserUseCase
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(usernaame string, password string) (string, error) {
	user, err := a.usecase.FindByUsernamePassword(usernaame, password)
	if err != nil {
		return "", fmt.Errorf("username of password not found")
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}

func GenerateToken(user model.UserCredential) (string, error) {
	return "token123", nil
}

func NewAuthUseCase(usercase UserUseCase) AuthUseCase {
	return &authUseCase{
		usecase: usercase,
	}
}
