package common

import "github.com/gin-gonic/gin"

type Handler struct {
	useCase IUseCase
	app     *gin.Engine
}

func NewHandler(app *gin.Engine, useCase IUseCase, prefixApi string) {
	handler := &Handler{
		app:     app,
		useCase: useCase,
	}

	handler.Routes(prefixApi)
}

func (h *Handler) Routes(prefix string) {
	_ = h.app.Group(prefix)
	{

	}
}
