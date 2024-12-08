package cart

import "github.com/google/uuid"

// Request for creating a cart
type CreateCartRequest struct {
	UserID uuid.UUID `json:"user_id"`
	// ProductID     uuid.UUID `json:"product_id"`
	TotalQuantity int     `json:"total_quantity"`
	TotalPrice    float64 `json:"total_price"`
}

// Request for updating a cart
type UpdateCartRequest struct {
	UserID        uuid.UUID `json:"user_id"`
	TotalQuantity int       `json:"total_quantity"`
	TotalPrice    float64   `json:"total_price"`
}

// Response for cart operations
type CartResponse struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	// ProductID     uuid.UUID `json:"product_id"`
	TotalQuantity int     `json:"total_quantity"`
	TotalPrice    float64 `json:"total_price"`
	// ProductName   string  `json:"product_name"`
}

// Response for deleting a cart
type DeleteCartResponse struct {
	Message string `json:"message"`
}

// Convert CartModel to CartResponse
func toCartResponse(cart *CartModel) *CartResponse {
	return &CartResponse{
		ID:     cart.ID,
		UserID: cart.UserID,
		// ProductID:     cart.ProductID,
		TotalQuantity: cart.TotalQuantity,
		TotalPrice:    cart.TotalPrice,
		// ProductName:   cart.Product.Name, // Assuming Product.Name is a field
	}
}
