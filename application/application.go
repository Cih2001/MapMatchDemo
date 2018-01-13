package application

import (
	"io"
	"text/template"

	"github.com/Cih2001/songo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Server      *echo.Echo = echo.New()
	apiPrefix   = "/api"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func setupStatics() {
	Server.Static("statics", "statics/")
}

func setupMiddlewares() {
	middlewares := []echo.MiddlewareFunc{
		middleware.Gzip(),
	}

	if Config.Debug {
		middlewares = append(middlewares, middleware.Logger())
	}

	Server.Use(middlewares...)
}

func StartServer() {
	songo.InitSongo(Config.MongoURL, Config.DatabaseName)

	t := &Template{
		templates: template.Must(template.ParseGlob(Config.Views + "*.html")),
	}
	Server.Renderer = t

	setupMiddlewares()
	setupStatics()
	setupRoutes()

	if err := Server.Start(Config.ListenAddress); err != nil {
		panic(err)
	}
}
