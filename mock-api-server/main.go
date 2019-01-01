package main

import (
	"github.com/infernalslam/mock-api-server/heath"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", heath.HeathCheck)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, heath.SayHello())
	// })
	// e.GET("/users/:id", heath.GetUser)
	e.Logger.Fatal(e.Start(":1323"))
}
