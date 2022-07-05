package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-trial/controllers"
)

const errorsSession = "errors"

type Template struct {
	templates *template.Template
}
  
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e := echo.New()
	// add ex method
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig {
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Debug = true
	e.Renderer = t

	e.GET("/", controllers.Home)
	e.PUT("/check", controllers.Check)
	e.POST("/check/list", controllers.ListCheck)
	
	e.Logger.Fatal(e.Start(":8080"))
}
