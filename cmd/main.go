package main

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"theprimeagen.tv/htmx/cmd/contact"
	"theprimeagen.tv/htmx/cmd/renderer"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.Renderer = renderer.NewTemplateRenderer()

	e.GET("/", func(c echo.Context) error {
		res := map[string]interface{}{
			"Name": "Per",
		}

    comp := hello("Test");
    comp.Render(context.Background(), os.Stdout)
		return c.Render(http.StatusOK, "index", res)
	})

  e.GET("/get-info", func(c echo.Context) error {
		res := map[string]interface{}{
			"Name": "Per",
      "Phone": "123",
      "Email": "p@gemail.com",
		}

    return c.Render(http.StatusOK, "name_card", res);
  })

	contact.CreateContactApi(e)

	e.Logger.Fatal(e.Start(":42069"))
}
