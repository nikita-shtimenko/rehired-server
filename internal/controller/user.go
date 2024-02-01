package controller

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/nikita-shtimenko/rehired-server/internal/usecase"
)

type userController struct {
	userUseCase *usecase.UserUseCase
	logger      *slog.Logger
}

func NewUserController(uc *usecase.UserUseCase, logger *slog.Logger) *userController {
	return &userController{userUseCase: uc, logger: logger}
}

func (uc *userController) GetAll(ctx echo.Context) error {
	users, err := uc.userUseCase.GetAll()

	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, users)
	return nil
}
