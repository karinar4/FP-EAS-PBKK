package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/configs"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"golang.org/x/crypto/bcrypt"
)

type IAuthUseCase interface {
	RegisterUser(*RegisterUserRequestDTO) e.ApiError
	LoginUser(*LoginUserRequestDTO) (*LoginUserResponseDTO, e.ApiError)
	GetMe(uuid.UUID) (*GetMeResponseDTO, e.ApiError)
	HashPassword(string) (string, error)
	VerifyPassword(string, string) bool
	GenerateToken(PayloadToken) (string, error)
	GetAllUser() (*GetAllUsersResponseDTO, e.ApiError)
	GetUserByEmail(string) (*GetUser, e.ApiError)
}

type authUseCase struct {
	authRepository IAuthRepository
}

func NewAuthUseCase(authRepository IAuthRepository) *authUseCase {
	return &authUseCase{
		authRepository: authRepository,
	}
}

func (uc *authUseCase) RegisterUser(data *RegisterUserRequestDTO) e.ApiError {
	userCheck, _ := uc.authRepository.GetUserByEmail(data.Email)

	if userCheck != nil {
		return e.NewApiError(400, "Email already registered")
	}

	hashedPassword, err := uc.HashPassword(data.Password)
	if err != nil {
		log.Println(err.Error())
		return e.NewApiError(500, fmt.Sprintf("Internal Server Error (%d)", e.ERROR_BCRYPT_HASH_FAILED))
	}

	user := &RegisterUserDomain{
		ID:       uuid.New(),
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
	}

	if err := uc.authRepository.RegisterUser(user); err != nil {
		log.Println(err.Error())
		return e.NewApiError(500, fmt.Sprintf("Internal Server Error (%d)", err.Code()))
	}
	return nil
}

func (uc *authUseCase) LoginUser(data *LoginUserRequestDTO) (*LoginUserResponseDTO, e.ApiError) {
	user, err := uc.authRepository.GetUserByEmail(data.Email)
	if err != nil {
		return nil, e.NewApiError(400, "User not found")
	}

	if !uc.VerifyPassword(user.Password, data.Password) {
		return nil, e.NewApiError(400, "Password is incorrect")
	}

	payloadToken := PayloadToken{
		ID:   user.ID,
		Role: user.Role,
	}

	token, errToken := uc.GenerateToken(payloadToken)
	fmt.Println(token)
	if errToken != nil {
		log.Println(errToken.Error())
		return nil, e.NewApiError(500, fmt.Sprintf("Internal Server Error (%d)", e.ERROR_GENERATE_TOKEN_FAILED))
	}
	
	return &LoginUserResponseDTO{
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}, nil
}

func (uc *authUseCase) GetMe(userID uuid.UUID) (*GetMeResponseDTO, e.ApiError) {
	user, err := uc.authRepository.GetUserByID(userID)
	if err != nil {
		return &GetMeResponseDTO{}, e.NewApiError(404, "User not found")
	}
	return &GetMeResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (uc *authUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (uc *authUseCase) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (uc *authUseCase) GenerateToken(payloadToken PayloadToken) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = payloadToken.ID
	claims["role"] = payloadToken.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	secretKey := configs.Config.JWT_SECRET
	if secretKey == "" {
		return "", errors.New("secret key for JWT is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (uc *authUseCase) GetAllUser() (*GetAllUsersResponseDTO, e.ApiError) {
	users, err := uc.authRepository.GetAllUser()
	if err != nil {
		return nil, e.NewApiError(500, fmt.Sprintf("Internal Server Error (%d)", err.Code()))
	}

	var response []GetUser
	for _, user := range users {
		response = append(response, GetUser{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &GetAllUsersResponseDTO{
		Users: response,
	}, nil
}

func (uc *authUseCase) GetUserByEmail(email string) (*GetUser, e.ApiError) {
	user, err := uc.authRepository.GetUserByEmail(email)
	if err != nil {
		if err.Code() == e.ERROR_USER_NOT_FOUND {
			return nil, e.NewApiError(404, "User not found")
		}
		return nil, e.NewApiError(500, fmt.Sprintf("Internal Server Error (%d)", err.Code()))
	}
	return &GetUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
