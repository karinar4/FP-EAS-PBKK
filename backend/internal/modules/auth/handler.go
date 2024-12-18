package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
	CustomValidator "github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/validator"
)

type AuthHandler struct {
	authUseCase IAuthUseCase
	app         *gin.Engine
}

func NewAuthHandler(app *gin.Engine, authUseCase IAuthUseCase, prefixApi string) {
	authHandler := &AuthHandler{
		app:         app,
		authUseCase: authUseCase,
	}

	authHandler.Routes(prefixApi)
}

func (ah *AuthHandler) Routes(prefix string) {
	authentication := ah.app.Group(prefix)
	{
		authentication.POST("/register", ah.Register)
		authentication.POST("/login", ah.Login)

		authentication.Use(middleware.AuthenticateJWT())
		{
			authentication.GET("/me", ah.GetMe)
			authentication.PUT("/:id", ah.UpdateUser)
			//authentication.GET("/email/:email", ah.GetUserByEmail)
		}
	}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var authentication RegisterUserRequestDTO
	if err := c.ShouldBindJSON(&authentication); err != nil {
		var errMessages = CustomValidator.FormatValidationErrors(err)
		c.JSON(400, app.NewErrorResponse("Validation Error", &errMessages))
		return
	}

	if err := ah.authUseCase.RegisterUser(&authentication); err != nil {
		errMsg := err.Error()
		c.JSON(err.Code(), app.NewErrorResponse("Failed to register user", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("User registered successfully", &RegisterUserResponseDTO{
		Name:  authentication.Name,
		Email: authentication.Email,
	}))
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var authentication LoginUserRequestDTO
	if err := c.ShouldBindJSON(&authentication); err != nil {
		var errMessages = CustomValidator.FormatValidationErrors(err)
		c.JSON(400, app.NewErrorResponse("Validation Error", &errMessages))
		return
	}

	token, err := ah.authUseCase.LoginUser(&authentication)
	if err != nil {
		errMsg := err.Error()
		c.JSON(err.Code(), app.NewErrorResponse("Failed to login user", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("User logged in successfully", token))
}

func (ah *AuthHandler) GetMe(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(400, app.NewErrorResponse("User ID not found in context", nil))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(400, app.NewErrorResponse("Invalid user ID type", nil))
		return
	}

	parsedID, errUuid := uuid.Parse(userIDStr)
	if errUuid != nil {
		c.JSON(400, app.NewErrorResponse("Invalid user ID", nil))
		return
	}

	user, err := ah.authUseCase.GetMe(parsedID)
	if err != nil {
		var errMsg = err.Error()
		c.JSON(err.Code(), app.NewErrorResponse("Failed to get user data", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("User data retrieved successfully", user))
}

func (ah *AuthHandler) GetAllUsers(c *gin.Context) {
	users, err := ah.authUseCase.GetAllUser()
	if err != nil {
		errMsg := err.Error()
		c.JSON(err.Code(), app.NewErrorResponse("Failed to get all users", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("All users retrieved successfully", users))
}

func (ah *AuthHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := ah.authUseCase.GetUserByEmail(email)
	if err != nil {
		errMsg := err.Error()
		c.JSON(err.Code(), app.NewErrorResponse("Failed to get user data", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("User data retrieved successfully", user))
}

func (h *AuthHandler) UpdateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.authUseCase.UpdateUser(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update user", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("User updated successfully", res))
}