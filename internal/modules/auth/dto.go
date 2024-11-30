package auth

type (
	LoginUserRequestDTO struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	LoginUserResponseDTO struct {
		Email string `json:"email"`
		Roles string `json:"roles"`
		Token string `json:"token"`
	}

	RegisterUserRequestDTO struct {
		Nama            string `json:"nama" binding:"required"`
		Email           string `json:"email" binding:"required,email"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	RegisterUserResponseDTO struct {
		Nama  string `json:"nama"`
		Email string `json:"email"`
	}

	GetMeResponseDTO struct {
		Nama  string `json:"nama"`
		Email string `json:"email"`
		Roles string `json:"roles"`
	}

	GetUser struct {
		ID    string `json:"id"`
		Nama  string `json:"nama"`
		Email string `json:"email"`
	}

	GetAllUsersResponseDTO struct {
		Users []GetUser `json:"users"`
	}
)
