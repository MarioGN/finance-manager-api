package usecase

import (
	"fmt"

	"github.com/MarioGN/finance-manager-api/internal/auth/dto"
	"github.com/MarioGN/finance-manager-api/internal/auth/entity"
	"github.com/MarioGN/finance-manager-api/internal/auth/repository"
)

func RegisterUser(r repository.UserRepository, input dto.RegisterUserDTO) (output *dto.RegisteredUserResponseDTO, err error) {
	newUser, err := entity.NewUserAccount(input.Email, input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create user account entity: %w", err)
	}

	id, err := r.Save(*newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to save user account: %w", err)
	}

	return &dto.RegisteredUserResponseDTO{
		ID:    id,
		Email: newUser.Email(),
	}, nil
}
