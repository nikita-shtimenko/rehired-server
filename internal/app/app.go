package app

import (
	"log"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nikita-shtimenko/rehired-server/internal/config"
	"github.com/nikita-shtimenko/rehired-server/internal/controller"
	"github.com/nikita-shtimenko/rehired-server/internal/repository"
	"github.com/nikita-shtimenko/rehired-server/internal/usecase"
	"github.com/nikita-shtimenko/rehired-server/pkg/logger"
	"github.com/nikita-shtimenko/rehired-server/pkg/postgres"
)

type Application struct {
	Config *config.Config
	Logger *slog.Logger
	DB     *pgxpool.Pool
}

func New() *Application {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	mylogger, err := logger.New(config.App.Env)
	if err != nil {
		log.Fatalf("logger error: %s", err)
	}

	db, err := postgres.New(config.Database.URL, time.Second*5)
	if err != nil {
		log.Fatalf("database error: %s", err)
	}

	application := &Application{
		Config: config,
		Logger: mylogger,
		DB:     db,
	}

	return application
}

func (a *Application) Run() {
	commonController := controller.NewCommonController(a.Logger)

	userRepository := repository.NewUserRepository(a.DB, a.Logger)
	userUseCase := usecase.NewUserUseCase(userRepository, a.Logger)
	userController := controller.NewUserController(userUseCase, a.Logger)

	echo := echo.New()
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	controller.BootstrapRoutes(echo, commonController, userController)

	echo.Logger.Fatal(echo.Start("localhost:" + a.Config.HTTP.Port))
}
