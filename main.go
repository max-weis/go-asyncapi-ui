package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
	"os"
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

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	asyncapi_url := os.Getenv("ASYNCAPI_URL")
	if asyncapi_url == "" {
		asyncapi_url = "http://localhost:8000/static/asyncapi.yml"
	}

	asyncapi_ui_url := os.Getenv("ASYNCAPI_UI_URL")
	if asyncapi_ui_url == "" {
		asyncapi_ui_url = "/asyncapi-ui"
	}

	e.Static("/static", ".")

	// Named route "foobar"
	e.GET(asyncapi_ui_url, func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"url": asyncapi_url,
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
