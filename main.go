// main.go
package main // import "github.com/go-module-test"

import (
	"net/http"

	"github.com/go-module-test/encode"
	"github.com/labstack/echo"
)

func main() {

	println(encode.EncodeB58([]byte("TEST")))

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, GO Moudules!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
