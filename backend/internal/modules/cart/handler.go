package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type CartHandler struct {
	cartUseCase ICartUseCase
	app         *gin.Engine
}

func NewCartHandler(app *gin.Engine, cartUseCase ICartUseCase, prefixApi string) {
	handler := &CartHandler{
		app:         app,
		cartUseCase: cartUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *CartHandler) Routes(prefix string) {
	cart := h.app.Group(prefix)
	cart.Use(middleware.AuthenticateJWT())
	{
		cart.POST("/", h.CreateCart)
		cart.GET("/", h.GetCartByUser)
		cart.PUT("/:id", h.UpdateCart)
		cart.DELETE("/:id", h.DeleteCart)
	}
}

func (h *CartHandler) CreateCart(c *gin.Context) {
	var req CreateCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	req.TotalQuantity = 0
	req.TotalPrice = 0

	res, err := h.cartUseCase.CreateCart(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create cart", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Cart created successfully", res))
}

func (h *CartHandler) GetCartByUser(c *gin.Context) {
	userID := c.GetString("user_id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid user ID", nil))
		return
	}

	res, err := h.cartUseCase.GetCartByUser(parsedUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve cart", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Cart retrieved successfully", &res))
}

func (h *CartHandler) UpdateCart(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	var req UpdateCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.cartUseCase.UpdateCart(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to update cart", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Cart updated successfully", res))
}

func (h *CartHandler) DeleteCart(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid ID format", nil))
		return
	}

	err = h.cartUseCase.DeleteCart(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete cart", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Cart deleted successfully", &DeleteCartResponse{
		Message: "Cart item deleted",
	}))
}
