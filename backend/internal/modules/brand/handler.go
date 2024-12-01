package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type BrandHandler struct {
	brandUseCase IBrandUseCase
	app         *gin.Engine
}

func NewBrandHandler(app *gin.Engine, brandUseCase IBrandUseCase, prefixApi string) {
	handler := &BrandHandler{
		app:         app,
		brandUseCase: brandUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *BrandHandler) Routes(prefix string) {
	brand := h.app.Group(prefix)
	brand.Use(middleware.AuthenticateJWT())
	{
		brand.POST("/", h.CreateBrand)
		brand.GET("/", h.GetAllBrand)
		brand.GET("/:id", h.GetBrandByID)
		brand.PUT("/:id", h.UpdateBrand)
		brand.DELETE("/:id", h.DeleteBrand)
	}
}

func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var req CreateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.brandUseCase.CreateBrand(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create brand", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Brand created successfully", res))
}

func (h *BrandHandler) GetAllBrand(c *gin.Context) {
	res, err := h.brandUseCase.GetAllBrand()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve brand list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Brand list retrieved successfully", res))
}

func (h *BrandHandler) GetBrandByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	res, err := h.brandUseCase.GetBrandByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Brand not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Brand retrieved successfully", res))
}

func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.brandUseCase.UpdateBrand(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update brand", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Brand updated successfully", res))
}

func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	if err := h.brandUseCase.DeleteBrand(id); err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete brand", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Brand deleted successfully", &DeleteBrandResponse{
		Message: "Brand has been deleted",
	}))
}
