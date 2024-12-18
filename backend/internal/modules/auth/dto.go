package auth

import "github.com/google/uuid"

type (
	LoginUserRequestDTO struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	LoginUserResponseDTO struct {
		Email string `json:"email"`
		Role  string `json:"role"`
		Token string `json:"token"`
	}

	RegisterUserRequestDTO struct {
		Name            string `json:"name" binding:"required"`
		Email           string `json:"email" binding:"required,email"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	RegisterUserResponseDTO struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	GetMeResponseDTO struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Role  string    `json:"role"`
		Telephone string `json:"telephone"`
		Address string `json:"address"`
	}

	UpdateUserRequest struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Telephone string `json:"telephone"`
		Address string `json:"address"`
	}

	UpdateUserResponse struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Telephone string `json:"telephone"`
		Address string `json:"address"`
	}

	GetUser struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
	}

	GetAllUsersResponseDTO struct {
		Users []GetUser `json:"users"`
	}
)
