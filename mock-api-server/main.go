package main

import (
	"net/http"

	"github.com/infernalslam/mock-api-server/heath"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, heath.SayHello())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
