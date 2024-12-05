package payment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type PaymentHandler struct {
	paymentUseCase IPaymentUseCase
	app         *gin.Engine
}

func NewPaymentHandler(app *gin.Engine, paymentUseCase IPaymentUseCase, prefixApi string) {
	handler := &PaymentHandler{
		app:         app,
		paymentUseCase: paymentUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *PaymentHandler) Routes(prefix string) {
	payment := h.app.Group(prefix)
	payment.Use(middleware.AuthenticateJWT())
	{
		payment.POST("/", h.CreatePayment)
		payment.GET("/", h.GetAllPayment)
		payment.GET("/:id", h.GetPaymentByID)
		payment.PUT("/:id", h.UpdatePayment)
		payment.DELETE("/:id", h.DeletePayment)
	}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.paymentUseCase.CreatePayment(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create payment", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Payment created successfully", res))
}

func (h *PaymentHandler) GetAllPayment(c *gin.Context) {
	res, err := h.paymentUseCase.GetAllPayment()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve payment list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Payment list retrieved successfully", res))
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	res, err := h.paymentUseCase.GetPaymentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Payment not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Payment retrieved successfully", res))
}

func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.paymentUseCase.UpdatePayment(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update payment", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Payment updated successfully", res))
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	if err := h.paymentUseCase.DeletePayment(id); err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete payment", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Payment deleted successfully", &DeletePaymentResponse{
		Message: "Payment has been deleted",
	}))
}
