package repository

import "github.com/MarioGN/finance-manager-api/internal/auth/entity"

type UserRepository interface {
	Save(user entity.UserAccount) (int64, error)
}
