package brand

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type IBrandRepository interface {
	CreateBrand(*BrandModel) (*BrandModel, e.ApiError)
	GetAllBrand() ([]BrandModel, e.ApiError)
	GetBrandByID(uuid.UUID) (*BrandModel, e.ApiError)
	UpdateBrand(*BrandModel) (*BrandModel, e.ApiError)
	DeleteBrand(uuid.UUID) e.ApiError
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *brandRepository {
	return &brandRepository{db}
}

func (r *brandRepository) CreateBrand(data *BrandModel) (*BrandModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *brandRepository) GetAllBrand() ([]BrandModel, e.ApiError) {
	var brands []BrandModel
	if err := r.db.Find(&brands).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return brands, nil
}

func (r *brandRepository) GetBrandByID(id uuid.UUID) (*BrandModel, e.ApiError) {
	brand := &BrandModel{}
	if err := r.db.Where("id = ?", id).First(brand).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewApiError(e.ErrNotFound, "Brand not found")
		}
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return brand, nil
}

func (r *brandRepository) UpdateBrand(data *BrandModel) (*BrandModel, e.ApiError) {
	if err := r.db.Save(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}
	return data, nil
}

func (r *brandRepository) DeleteBrand(id uuid.UUID) e.ApiError {
	result := r.db.Delete(&BrandModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Brand not found")
	}
	return nil
}
