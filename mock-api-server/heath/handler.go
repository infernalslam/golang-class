package heath

import (
	"net/http"

	"github.com/labstack/echo"
)

// HeathCheck controller
func HeathCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Hello welcome")
}

// GetUser controller
func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
