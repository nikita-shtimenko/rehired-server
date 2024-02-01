package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikita-shtimenko/rehired-server/internal/entity"
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger *slog.Logger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (r *UserRepository) GetAll() ([]entity.User, error) {
	users := []entity.User{
		{
			Name:  "Test1",
			Email: "test1@gmail.com",
		},
		{
			Name:  "Test2",
			Email: "test2@gmail.com",
		},
		{
			Name:  "Test3",
			Email: "test3@gmail.com",
		},
	}

	return users, nil
}

func (r *UserRepository) GetById(id int) (entity.User, error) {
	user := entity.User{Name: "Test", Email: "test@gmail.com"}
	return user, nil
}
