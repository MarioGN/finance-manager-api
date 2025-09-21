package entity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserAccount struct {
	id           int64
	email        string
	passwordHash string
}

func NewUserAccount(email, password string) (*UserAccount, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	if len(password) < 6 {
		return nil, fmt.Errorf("password must be at least 6 characters long")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return &UserAccount{
		email:        email,
		passwordHash: string(pw),
	}, nil
}

func (u *UserAccount) ID() int64 {
	return u.id
}

func (u *UserAccount) Email() string {
	return u.email
}

func (u *UserAccount) PasswordHash() string {
	return u.passwordHash
}

func (u *UserAccount) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.passwordHash), []byte(password))
}
