package merk

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type MerkHandler struct {
	merkUseCase IMerkUseCase
	app         *gin.Engine
}

func NewMerkHandler(app *gin.Engine, merkUseCase IMerkUseCase, prefixApi string) {
	handler := &MerkHandler{
		app:         app,
		merkUseCase: merkUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *MerkHandler) Routes(prefix string) {
	merk := h.app.Group(prefix)
	merk.Use(middleware.AuthenticateJWT())
	{
		merk.POST("/", h.CreateMerk)
		merk.GET("/", h.GetAllMerk)
		merk.GET("/:id", h.GetMerkByID)
		merk.PUT("/:id", h.UpdateMerk)
		merk.DELETE("/:id", h.DeleteMerk)
	}
}

func (h *MerkHandler) CreateMerk(c *gin.Context) {
	var req CreateMerkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.merkUseCase.CreateMerk(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create merk", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Merk created successfully", res))
}

func (h *MerkHandler) GetAllMerk(c *gin.Context) {
	res, err := h.merkUseCase.GetAllMerk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve merk list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Merk list retrieved successfully", res))
}

func (h *MerkHandler) GetMerkByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	res, err := h.merkUseCase.GetMerkByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Merk not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Merk retrieved successfully", res))
}

func (h *MerkHandler) UpdateMerk(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdateMerkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.merkUseCase.UpdateMerk(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update merk", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Merk updated successfully", res))
}

func (h *MerkHandler) DeleteMerk(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	if err := h.merkUseCase.DeleteMerk(id); err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete merk", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Merk deleted successfully", &DeleteMerkResponse{
		Message: "Merk has been deleted",
	}))
}
