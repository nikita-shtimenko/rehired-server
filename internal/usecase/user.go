package usecase

import (
	"log/slog"

	"github.com/nikita-shtimenko/rehired-server/internal/entity"
	"github.com/nikita-shtimenko/rehired-server/internal/repository"
)

type UserUseCase struct {
	userRepository *repository.UserRepository
	logger         *slog.Logger
}

func NewUserUseCase(userRepo *repository.UserRepository, logger *slog.Logger) *UserUseCase {
	return &UserUseCase{userRepository: userRepo, logger: logger}
}

func (uc *UserUseCase) GetAll() ([]entity.User, error) {
	return uc.userRepository.GetAll()
}
