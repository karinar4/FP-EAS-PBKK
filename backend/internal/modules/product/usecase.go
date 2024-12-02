package product

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/brand"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/category"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type IProductUseCase interface {
	CreateProduct(*CreateProductRequest) (*CreateProductResponse, e.ApiError)
	GetAllProducts() (*GetAllProductsResponse, e.ApiError)
	GetProductByID(uuid.UUID) (*GetProductResponse, e.ApiError)
	UpdateProduct(uuid.UUID, *UpdateProductRequest) (*UpdateProductResponse, e.ApiError)
	DeleteProduct(uuid.UUID) e.ApiError
}

type productUseCase struct {
	repo IProductRepository
}

func NewProductUseCase(repo IProductRepository) IProductUseCase {
	return &productUseCase{repo: repo}
}

func (uc *productUseCase) CreateProduct(req *CreateProductRequest) (*CreateProductResponse, e.ApiError) {
	product := &ProductModel{
		BaseModels: common.BaseModels{
			ID: uuid.New(),
		},
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		BrandID:     req.BrandID,
	}

	result, err := uc.repo.CreateProduct(product)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateProductResponse{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		Stock:       result.Stock,
	}, nil
}

func (uc *productUseCase) GetAllProducts() (*GetAllProductsResponse, e.ApiError) {
	products, err := uc.repo.GetAllProducts()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetProductResponse
	for _, product := range products {
		response = append(response, GetProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			//CategoryID:  product.CategoryID,
			//BrandID:     product.BrandID,
			Category: category.GetCategoryResponse{
				ID:   product.Category.ID,
				Name: product.Category.Name,
			},
			Brand: brand.GetBrandResponse{
				ID:   product.Brand.ID,
				Name: product.Brand.Name,
			},
		})
	}

	return &GetAllProductsResponse{Products: response}, nil
}

func (uc *productUseCase) GetProductByID(id uuid.UUID) (*GetProductResponse, e.ApiError) {
	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return &GetProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		//CategoryID:  product.CategoryID,
		//BrandID:     product.BrandID,
		Category: category.GetCategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
		Brand: brand.GetBrandResponse{
			ID:   product.Brand.ID,
			Name: product.Brand.Name,
		},
	}, nil
}

func (uc *productUseCase) UpdateProduct(id uuid.UUID, req *UpdateProductRequest) (*UpdateProductResponse, e.ApiError) {
	updatedFields := make(map[string]interface{})
	if req.Name != "" {
		updatedFields["name"] = req.Name
	}
	if req.Description != "" {
		updatedFields["description"] = req.Description
	}
	if req.Price != 0 {
		updatedFields["price"] = req.Price
	}
	if req.Stock != 0 {
		updatedFields["stock"] = req.Stock
	}
	if req.CategoryID != uuid.Nil {
		updatedFields["category_id"] = req.CategoryID
	}
	if req.BrandID != uuid.Nil {
		updatedFields["brand_id"] = req.BrandID
	}

	updatedProduct, err := uc.repo.UpdateProduct(id, updatedFields)
	if err != nil {
		return nil, err
	}

	return &UpdateProductResponse{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		Stock:       updatedProduct.Stock,
		CategoryID:  updatedProduct.CategoryID,
		BrandID:     updatedProduct.BrandID,
	}, nil
}

func (uc *productUseCase) DeleteProduct(id uuid.UUID) e.ApiError {
	return uc.repo.DeleteProduct(id)
}
