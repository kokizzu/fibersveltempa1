package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/Z"
)

//go:embed myproject1/index.html
var indexHtml string

func main() {
	app := fiber.New()

	users := []struct {
		Name string
		Age  int
	}{
		{
			`John`,
			20,
		},
		{
			`Jane`,
			30,
		},
	}

	// when using this method, when deploying, you must rsync also the html
	//file := "myproject1/index.html"
	//tc, err := Z.ParseFile(true, true, file)
	//L.PanicIf(err, `Z.ParseFile `+file)

	// when using this method, when deploying, only need to copy the binary only
	tc := Z.FromString(indexHtml, true)
	app.
		Get("/", func(c *fiber.Ctx) error {
			b := bytes.Buffer{}
			tc.Render(&b, M.SX{
				`title`: `test`,
				`obj1`: M.SX{
					`a`: 1,
					`b`: 2,
				},
				`arr`:  users,
				`str1`: `something`,
			})

			c.Set(`Content-Type`, `text/html`)
			return c.SendString(b.String())
		})

	log.Fatal(app.Listen(":3001"))
}
