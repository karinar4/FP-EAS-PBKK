package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type CategoryHandler struct {
	categoryUseCase ICategoryUseCase
	app             *gin.Engine
}

func NewCategoryHandler(app *gin.Engine, categoryUseCase ICategoryUseCase, prefixApi string) {
	handler := &CategoryHandler{
		app:             app,
		categoryUseCase: categoryUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *CategoryHandler) Routes(prefix string) {
	category := h.app.Group(prefix)

	category.GET("/", h.GetAllCategory)
	category.GET("/:id", h.GetCategoryByID)

	category.Use(middleware.AuthenticateJWT())
	{
		category.POST("/", h.CreateCategory)
		category.PUT("/:id", h.UpdateCategory)
		category.DELETE("/:id", h.DeleteCategory)
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.categoryUseCase.CreateCategory(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create category", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Category created successfully", res))
}

func (h *CategoryHandler) GetAllCategory(c *gin.Context) {
	res, err := h.categoryUseCase.GetAllCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve category list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Category list retrieved successfully", res))
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	res, err := h.categoryUseCase.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Category not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Category retrieved successfully", res))
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.categoryUseCase.UpdateCategory(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update category", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Category updated successfully", res))
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	if err := h.categoryUseCase.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete category", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Category deleted successfully", &DeleteCategoryResponse{
		Message: "Category has been deleted",
	}))
}
