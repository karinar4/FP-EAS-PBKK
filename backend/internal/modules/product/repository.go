package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(*ProductModel) (*ProductModel, e.ApiError)
	GetAllProducts() ([]ProductModel, e.ApiError)
	GetProductByID(uuid.UUID) (*ProductModel, e.ApiError)
	UpdateProduct(id uuid.UUID, updatedFields map[string]interface{}) (*ProductModel, e.ApiError)
	DeleteProduct(uuid.UUID) e.ApiError
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(data *ProductModel) (*ProductModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *productRepository) GetAllProducts() ([]ProductModel, e.ApiError) {
	var products []ProductModel
	if err := r.db.Preload("Category").Preload("Brand").Preload("Images").Find(&products).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return products, nil
}

func (r *productRepository) GetProductByID(id uuid.UUID) (*ProductModel, e.ApiError) {
	var product ProductModel
	if err := r.db.Preload("Category").Preload("Brand").Preload("Images").First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewApiError(e.ErrNotFound, "Product not found")
		}
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return &product, nil
}

func (r *productRepository) UpdateProduct(id uuid.UUID, updatedFields map[string]interface{}) (*ProductModel, e.ApiError) {
	if err := r.db.Model(&ProductModel{}).Where("id = ?", id).Updates(updatedFields).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}
	var updatedProduct ProductModel
	if err := r.db.Preload("Category").Preload("Brand").Preload("Images").First(&updatedProduct, "id = ?", id).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return &updatedProduct, nil
}

func (r *productRepository) DeleteProduct(id uuid.UUID) e.ApiError {
	result := r.db.Delete(&ProductModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Product not found")
	}
	return nil
}
