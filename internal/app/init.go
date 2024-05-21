package app

import (
	"errors"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
)

var (
	app  *Application
	once sync.Once
)

func InitApplication() (*Application, error) {
	if app == nil {
		once.Do(func() {
			err := setApplication()
			if err != nil {
				log.Fatal(err)
			}
		})
	} else {
		return nil, errors.New("something went wrong while configuaring application")
	}
	return app, nil
}

func (app *Application) Run() error {
	app.Handler.Start(":3000")
	return nil
}

func setApplication() error {
	handler := echo.New()
	app = &Application{
		Handler: handler,
	}
	app.InitRoutes()
	return nil
}
