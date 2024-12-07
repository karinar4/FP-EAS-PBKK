package image

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/app"
)

type ImageHandler struct {
	imageUseCase IImageUseCase
}

func NewImageHandler(router *gin.Engine, useCase IImageUseCase, prefix string) {
	handler := &ImageHandler{imageUseCase: useCase}

	imageGroup := router.Group(prefix)
	{
		imageGroup.POST("/", handler.CreateImage)
		imageGroup.GET("/", handler.GetAllImages)    // Get all images
		imageGroup.GET("/:id", handler.GetImageByID) // Get image by ID
		imageGroup.GET("/product/:product_id", handler.GetImagesByProductID)
		imageGroup.DELETE("/:id", handler.DeleteImage)
	}
}

func (h *ImageHandler) CreateImage(c *gin.Context) {
	var req CreateImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid request payload", nil))
		return
	}

	res, err := h.imageUseCase.CreateImage(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to create image", nil))
		return
	}

	c.JSON(http.StatusCreated, app.NewSuccessResponse("Image created successfully", res))
}

func (h *ImageHandler) GetAllImages(c *gin.Context) {
	res, err := h.imageUseCase.GetAllImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve images", nil))
		return
	}

	response := struct {
		Images []GetImageResponse `json:"images"`
	}{
		Images: res,
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Images retrieved successfully", &response))
}

func (h *ImageHandler) GetImageByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.imageUseCase.GetImageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, app.NewErrorResponse("Image not found", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Image retrieved successfully", res))
}

func (h *ImageHandler) GetImagesByProductID(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid Product ID format", nil))
		return
	}

	res, err := h.imageUseCase.GetImagesByProductID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to retrieve images", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Images retrieved successfully", res))
}

func (h *ImageHandler) DeleteImage(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, app.NewErrorResponse("Invalid Image ID format", nil))
		return
	}

	err = h.imageUseCase.DeleteImage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.NewErrorResponse("Failed to delete image", nil))
		return
	}

	c.JSON(http.StatusOK, app.NewSuccessResponse("Image deleted successfully", &DeleteImageResponse{
		Message: "Image has been deleted",
	}))

}
