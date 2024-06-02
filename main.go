package main

import (
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/mendoncas/tigrinhobet/models"
)

type Templates struct {
	templates *template.Template
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Count struct {
	Count int
}

type Params struct {
	Bets *[]models.Bet
}

func main() {
	e := echo.New()
	bets := make([]models.Bet, 0)
	bets = append(
		bets,
		models.Bet{Value: "200", Description: "vai dar merda", BetterEmail: "caio@unifor.br"},
		models.Bet{Value: "2.50", Description: "vai dar bom", BetterEmail: "pedro@unifor.br"},
		models.Bet{Value: "999", Description: "vai ser o laio", BetterEmail: "ripardo@unifor.br"},
	)
	params := Params{Bets: &bets}
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	count := Count{Count: 0}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", params)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	})

	e.POST("/bet", func(c echo.Context) error {
		bets = append(bets, models.Bet{
			Description: c.FormValue("description"),
			BetterEmail: c.FormValue("email"),
			Value:       c.FormValue("value"),
		})
		return c.Render(200, "betDisplay", params)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
