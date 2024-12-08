package cart_product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type CartProductHandler struct {
	cartProductUseCase ICartProductUseCase
	app                *gin.Engine
}

func NewCartProductHandler(app *gin.Engine, cartProductUseCase ICartProductUseCase, prefixApi string) {
	handler := &CartProductHandler{
		app:                app,
		cartProductUseCase: cartProductUseCase,
	}

	handler.Routes(prefixApi)
}

func (h *CartProductHandler) Routes(prefix string) {
	cart_product := h.app.Group(prefix)
	cart_product.Use(middleware.AuthenticateJWT())
	{
		cart_product.POST("/", h.CreateCartProduct)
		cart_product.GET("/", h.GetAllCartProducts)
		cart_product.DELETE("/:cart_id/:product_id", h.DeleteCartProduct)
	}
}

func (h *CartProductHandler) CreateCartProduct(c *gin.Context) {
	var req CreateCartProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.cartProductUseCase.CreateCartProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create cart-product", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Cart-product created successfully", res))
}

func (h *CartProductHandler) GetAllCartProducts(c *gin.Context) {
	res, err := h.cartProductUseCase.GetAllCartProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve cart-product list", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Cart-product list retrieved successfully", res))
}

func (h *CartProductHandler) DeleteCartProduct(c *gin.Context) {
	cartID, err := uuid.Parse(c.Param("cart_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid cart ID format", nil))
		return
	}

	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid product ID format", nil))
		return
	}

	err = h.cartProductUseCase.DeleteCartProduct(cartID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete cart-product", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Cart-product deleted successfully", &DeleteCartProductResponse{
		Message: "Cart-product has been deleted",
	}))
}
