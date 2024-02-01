package controller

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type commonController struct {
	logger *slog.Logger
}

func NewCommonController(logger *slog.Logger) *commonController {
	return &commonController{logger: logger}
}

func (cc *commonController) homepage(ctx echo.Context) error {
	ctx.String(http.StatusOK, "homepage")
	return nil
}
