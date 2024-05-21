package app

import (
	home "github.com/MJU-Capstone-6/devmark-frontend/static/html"
	"github.com/labstack/echo/v4"
)

func (app *Application) InitRoutes() {
	app.Handler.Static("/static", "static")
	app.InitHomeRoutes()
}

func (app *Application) InitHomeRoutes() {
	e := app.Handler.Group("/home")
	e.GET("", func(ctx echo.Context) error {
		homeComponent := home.Hello("milky")
		return homeComponent.Render(ctx.Request().Context(), ctx.Response().Writer)
	})
}
