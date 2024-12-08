package product_transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type ProductTransactionHandler struct {
	productTransactionUseCase IProductTransactionUseCase
	app            *gin.Engine
}

func NewProductTransactionHandler(app *gin.Engine, productTransactionUseCase IProductTransactionUseCase, prefixApi string) {
	handler := &ProductTransactionHandler{
		app:            app,
		productTransactionUseCase: productTransactionUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *ProductTransactionHandler) Routes(prefix string) {
	product_transaction := h.app.Group(prefix)
	product_transaction.Use(middleware.AuthenticateJWT())
	{
		product_transaction.POST("/", h.CreateProductTransaction)
		product_transaction.GET("/", h.GetAllProductTransactions)
		// product_transaction.GET("/:transaction_id", h.GetByTransactionID)
		// product_transaction.PUT("/:id", h.UpdateProduct)
		product_transaction.DELETE("/:transaction_id/:product_id", h.DeleteProductTransaction)
	}
}

func (h *ProductTransactionHandler) CreateProductTransaction(c *gin.Context) {
	var req CreateProductTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.productTransactionUseCase.CreateProductTransaction(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create product-transaction", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Product-transaction created successfully", res))
}

func (h *ProductTransactionHandler) GetAllProductTransactions(c *gin.Context) {
	res, err := h.productTransactionUseCase.GetAllProductTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve product-transaction list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Product-transaction list retrieved successfully", res))
}

// func (h *ProductTransactionHandler) GetByTransactionID(c *gin.Context) {
// 	transactionID, err := uuid.Parse(c.Param("transaction_id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
// 		return
// 	}

// 	res, err := h.productTransactionUseCase.GetByTransactionID(transactionID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, app.NewErrorResponse("Product not found", nil))
// 		return
// 	}

// 	c.JSON(http.StatusOK, app.NewSuccessResponse("Product retrieved successfully", res))
// }

// func (h *ProductTransactionHandler) UpdateProduct(c *gin.Context) {
// 	id, err := uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
// 		return
// 	}

// 	var req UpdateProductRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
// 		return
// 	}

// 	res, err := h.productUseCase.UpdateProduct(id, &req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update product", nil))
// 		return
// 	}

// 	c.JSON(http.StatusOK, app.NewSuccessResponse("Product updated successfully", res))
// }

func (h *ProductTransactionHandler) DeleteProductTransaction(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	transactionID, err := uuid.Parse(c.Param("transaction_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	err = h.productTransactionUseCase.DeleteProductTransaction(productID, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete product-transaction", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Product deleted successfully", &DeleteProductTransactionResponse{
		Message: "Product-transaction has been deleted",
	}))
}
