package transaction

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type TransactionHandler struct {
	transactionUseCase ITransactionUseCase
	app         *gin.Engine
}

func NewTransactionHandler(app *gin.Engine, transactionUseCase ITransactionUseCase, prefixApi string) {
	handler := &TransactionHandler{
		app:         app,
		transactionUseCase: transactionUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *TransactionHandler) Routes(prefix string) {
	transaction := h.app.Group(prefix)
	transaction.Use(middleware.AuthenticateJWT())
	{
		transaction.POST("/", h.CreateTransaction)
		transaction.GET("/", h.GetAllTransaction)
		transaction.GET("/:id", h.GetTransactionByID)
		transaction.GET("/user/:user_id", h.GetAllTransactionByUserID)
		transaction.PUT("/:id", h.UpdateTransaction)
		transaction.DELETE("/:id", h.DeleteTransaction)
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.transactionUseCase.CreateTransaction(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create transaction", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Transaction created successfully", res))
}

func (h *TransactionHandler) GetAllTransaction(c *gin.Context) {
	res, err := h.transactionUseCase.GetAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve transaction list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Transaction list retrieved successfully", res))
}

func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	res, err := h.transactionUseCase.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Transaction not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Transaction retrieved successfully", res))
}

func (h *TransactionHandler) GetAllTransactionByUserID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	res, err := h.transactionUseCase.GetAllTransactionByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Transaction not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Transaction retrieved successfully", res))
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.transactionUseCase.UpdateTransaction(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update transaction", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Transaction updated successfully", res))
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	if err := h.transactionUseCase.DeleteTransaction(id); err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete transaction", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Transaction deleted successfully", &DeleteTransactionResponse{
		Message: "Transaction has been deleted",
	}))
}
