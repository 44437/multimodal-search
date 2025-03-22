package multimodal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetTest(ec echo.Context) error
	RegisterRoutes(e *echo.Echo)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/test", h.GetTest)
}

func (h *handler) GetTest(ctx echo.Context) error {
	err := h.service.GetTest(ctx.Request().Context())
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"dfsdf": "dfsdf",
	})
}
