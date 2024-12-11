package cart_product

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/product"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type ICartProductUseCase interface {
	CreateCartProduct(*CreateCartProductRequest) (*CreateCartProductResponse, e.ApiError)
	GetAllCartProducts() (*GetAllCartProductsResponse, e.ApiError)
	GetCartProductsByID(uuid.UUID) (*GetAllCartProductsResponse, e.ApiError)
	UpdateCartProduct(uuid.UUID, uuid.UUID, *UpdateCartProductRequest) (*UpdateCartProductResponse, e.ApiError)
	DeleteCartProduct(uuid.UUID, uuid.UUID) e.ApiError
}

type cartProductUseCase struct {
	repo ICartProductRepository
}

func NewCartProductUseCase(repo ICartProductRepository) ICartProductUseCase {
	return &cartProductUseCase{repo: repo}
}

func (uc *cartProductUseCase) CreateCartProduct(req *CreateCartProductRequest) (*CreateCartProductResponse, e.ApiError) {
	cartProduct := &CartProductModel{
		CartID:        req.CartID,
		ProductID:     req.ProductID,
		RentStartDate: req.RentStartDate,
		RentEndDate:   req.RentEndDate,
		Quantity:      req.Quantity,
		Price:         req.Price,
	}

	result, err := uc.repo.CreateCartProduct(cartProduct)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateCartProductResponse{
		CartID:        result.CartID,
		ProductID:     result.ProductID,
		RentStartDate: result.RentStartDate,
		RentEndDate:   result.RentEndDate,
		Quantity:      result.Quantity,
		Price:         result.Price,
	}, nil
}

func (uc *cartProductUseCase) GetAllCartProducts() (*GetAllCartProductsResponse, e.ApiError) {
	cartProducts, err := uc.repo.GetAllCartProducts()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetCartProductResponse
	for _, cartProduct := range cartProducts {
		response = append(response, GetCartProductResponse{
			CartID:        cartProduct.CartID,
			Product: product.GetProductResponse{
				ID:   cartProduct.Product.ID,
				Name: cartProduct.Product.Name,
				Price: cartProduct.Product.Price,
				Stock: cartProduct.Product.Stock,
			},
			RentStartDate: cartProduct.RentStartDate,
			RentEndDate:   cartProduct.RentEndDate,
			Quantity:      cartProduct.Quantity,
			Price:         cartProduct.Price,
		})
	}

	return &GetAllCartProductsResponse{CartProducts: response}, nil
}

func (uc *cartProductUseCase) GetCartProductsByID(cart_id uuid.UUID) (*GetAllCartProductsResponse, e.ApiError) {
	cartProducts, err := uc.repo.GetCartProductsByID(cart_id)
	if err != nil {
		return nil, err
	}

	var response []GetCartProductResponse
	for _, cartProduct := range cartProducts {
		response = append(response, GetCartProductResponse{
			CartID:        cartProduct.CartID,
			Product: product.GetProductResponse{
				ID:   cartProduct.Product.ID,
				Name: cartProduct.Product.Name,
				Price: cartProduct.Product.Price,
				Stock: cartProduct.Product.Stock,
			},
			RentStartDate: cartProduct.RentStartDate,
			RentEndDate:   cartProduct.RentEndDate,
			Quantity:      cartProduct.Quantity,
			Price:         cartProduct.Price,
		})
	}

	return &GetAllCartProductsResponse{CartProducts: response}, nil
}

func (uc *cartProductUseCase) UpdateCartProduct(cart_id uuid.UUID, product_id uuid.UUID, req *UpdateCartProductRequest) (*UpdateCartProductResponse, e.ApiError) {
	updatedFields := make(map[string]interface{})
	if !req.RentStartDate.IsZero() {
		updatedFields["rent_start_date"] = req.RentStartDate
	}
	if !req.RentEndDate.IsZero() {
		updatedFields["rent_end_date"] = req.RentEndDate
	}
	if req.Quantity != 0 {
		updatedFields["quantity"] = req.Quantity
	}
	if req.Price != 0 {
		updatedFields["price"] = req.Price
	}

	updatedCartProduct, err := uc.repo.UpdateCartProduct(cart_id, product_id, updatedFields)
	if err != nil {
		return nil, err
	}

	return &UpdateCartProductResponse{
		CartID:          updatedCartProduct.CartID,
		ProductID:        updatedCartProduct.ProductID,
		RentStartDate: updatedCartProduct.RentStartDate,
		RentEndDate:       updatedCartProduct.RentEndDate,
		Quantity:       updatedCartProduct.Quantity,
		Price:  updatedCartProduct.Price,
	}, nil
}

func (uc *cartProductUseCase) DeleteCartProduct(cartID uuid.UUID, productID uuid.UUID) e.ApiError {
	return uc.repo.DeleteCartProduct(cartID, productID)
}
