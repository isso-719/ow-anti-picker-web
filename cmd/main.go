package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"ow-anti-picker-web/src"
)

func main() {
	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func initRouting(e *echo.Echo) {
	initTemplate(e)
	e.GET("/", src.Index)

	apiGroup := e.Group("/api")
	apiGroup.POST("/get-anti-pick", src.GetAntiPick)

	e.Static("/public", "public")
}

func initTemplate(e *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
}
