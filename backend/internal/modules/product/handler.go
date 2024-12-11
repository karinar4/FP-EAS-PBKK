package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type ProductHandler struct {
	productUseCase IProductUseCase
	app            *gin.Engine
}

func NewProductHandler(app *gin.Engine, productUseCase IProductUseCase, prefixApi string) {
	handler := &ProductHandler{
		app:            app,
		productUseCase: productUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *ProductHandler) Routes(prefix string) {
	product := h.app.Group(prefix)

	product.GET("/", h.GetAllProducts)
	product.GET("/:id", h.GetProductByID)

	product.Use(middleware.AuthenticateJWT())
	{
		product.POST("/", h.CreateProduct)
		product.PUT("/:id", h.UpdateProduct)
		product.DELETE("/:id", h.DeleteProduct)
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.productUseCase.CreateProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create product", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Product created successfully", res))
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	res, err := h.productUseCase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve product list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Product list retrieved successfully", res))
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	res, err := h.productUseCase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Product not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Product retrieved successfully", res))
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.productUseCase.UpdateProduct(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update product", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Product updated successfully", res))
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	err = h.productUseCase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete product", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Product deleted successfully", &DeleteProductResponse{
		Message: "Product has been deleted",
	}))
}
