package dto

type RegisterUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisteredUserResponseDTO struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
