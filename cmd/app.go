package main

import (
	"errors"
	"net/http"

	"github.com/jeff-moorhead-pline/users-service/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	e *echo.Echo
}

func NewApp() *App {

	e := api.AttachHandlers(echo.New())

	return &App{e: e}
}

func (a *App) Run() error {

	a.e.Use(middleware.Logger())

	if err := a.e.Start(":8080"); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
