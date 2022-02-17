package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Static("/", ".")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/asyncapi-ui", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"url": "http://localhost:8000/asyncapi.yml",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
