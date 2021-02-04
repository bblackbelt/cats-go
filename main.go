package main

import (
	"net/http"

	"cats/domain"
	"cats/domain/adapter"
	ihttp "cats/executor"
	"cats/middleware"
	"cats/repo"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	executor := ihttp.NewHttpExecutor(ihttp.NewHttpClient(&http.Client{}))
	r := repo.NewCatRepository(executor, ihttp.NewHttpRequestFactory())

	usecase := domain.NewCatUseCase(r, adapter.NewCatMapper())
	middleware.RegisterCatsHandler(e, usecase)

	e.Start(":9090")
}
