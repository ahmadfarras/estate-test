package main

import (
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/usecase"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var (
		repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
			Dsn: dbDsn,
		})

		estateUsecase = usecase.NewEstateUsecase(repo)
	)
	opts := handler.NewServerOptions{
		EstateUsecase: estateUsecase,
		Validator:     validator.New(),
	}

	return handler.NewServer(opts)
}
