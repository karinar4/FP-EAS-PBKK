package image

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type IImageUseCase interface {
	CreateImage(*CreateImageRequest) (*CreateImageResponse, e.ApiError)
	GetAllImages() ([]GetImageResponse, e.ApiError)
	GetImageByID(string) (*GetImageResponse, e.ApiError)
	GetImagesByProductID(uuid.UUID) (*GetAllImagesResponse, e.ApiError)
	DeleteImage(uuid.UUID) e.ApiError
}

type imageUseCase struct {
	repo IImageRepository
}

func NewImageUseCase(repo IImageRepository) IImageUseCase {
	return &imageUseCase{repo: repo}
}

func (uc *imageUseCase) CreateImage(req *CreateImageRequest) (*CreateImageResponse, e.ApiError) {
	image := &ImageModel{
		BaseModels: common.BaseModels{
			ID: uuid.New(),
		},
		ProductID: req.ProductID,
		URL:       req.URL,
	}

	result, err := uc.repo.CreateImage(image)
	if err != nil {
		return nil, err
	}

	return &CreateImageResponse{
		ID:        result.ID,
		ProductID: result.ProductID,
		URL:       result.URL,
	}, nil
}

func (uc *imageUseCase) GetAllImages() ([]GetImageResponse, e.ApiError) {
	images, err := uc.repo.GetAllImages()
	if err != nil {
		return nil, err
	}

	var response []GetImageResponse
	for _, image := range images {
		response = append(response, GetImageResponse{
			ID:        image.ID,
			ProductID: image.ProductID,
			URL:       image.URL,
		})
	}

	return response, nil
}

func (uc *imageUseCase) GetImageByID(id string) (*GetImageResponse, e.ApiError) {
	image, err := uc.repo.GetImageByID(id)
	if err != nil {
		return nil, err
	}

	return &GetImageResponse{
		ID:        image.ID,
		ProductID: image.ProductID,
		URL:       image.URL,
	}, nil
}

func (uc *imageUseCase) GetImagesByProductID(productID uuid.UUID) (*GetAllImagesResponse, e.ApiError) {
	images, err := uc.repo.GetImagesByProductID(productID)
	if err != nil {
		return nil, err
	}

	var response []GetImageResponse
	for _, img := range images {
		response = append(response, GetImageResponse{
			ID:        img.ID,
			ProductID: img.ProductID,
			URL:       img.URL,
		})
	}

	return &GetAllImagesResponse{Images: response}, nil
}

func (uc *imageUseCase) DeleteImage(id uuid.UUID) e.ApiError {
	return uc.repo.DeleteImage(id)
}
