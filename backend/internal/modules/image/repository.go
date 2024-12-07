package image

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type IImageRepository interface {
	CreateImage(*ImageModel) (*ImageModel, e.ApiError)
	GetAllImages() ([]ImageModel, e.ApiError)
	GetImageByID(string) (*ImageModel, e.ApiError)
	GetImagesByProductID(uuid.UUID) ([]ImageModel, e.ApiError)
	DeleteImage(uuid.UUID) e.ApiError
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *imageRepository {
	return &imageRepository{db: db}
}

func (r *imageRepository) CreateImage(image *ImageModel) (*ImageModel, e.ApiError) {
	if err := r.db.Create(image).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return image, nil
}

func (r *imageRepository) GetAllImages() ([]ImageModel, e.ApiError) {
	var images []ImageModel
	if err := r.db.Find(&images).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return images, nil
}

func (r *imageRepository) GetImageByID(id string) (*ImageModel, e.ApiError) {
	var image ImageModel
	if err := r.db.First(&image, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewApiError(e.ErrNotFound, "Image not found")
		}
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return &image, nil
}

func (r *imageRepository) GetImagesByProductID(productID uuid.UUID) ([]ImageModel, e.ApiError) {
	var images []ImageModel
	if err := r.db.Where("product_id = ?", productID).Find(&images).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return images, nil
}

func (r *imageRepository) DeleteImage(id uuid.UUID) e.ApiError {
	if err := r.db.Delete(&ImageModel{}, id).Error; err != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, err.Error())
	}
	return nil
}
