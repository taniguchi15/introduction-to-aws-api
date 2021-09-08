package api

import (
	"github.com/labstack/echo"
)

func log(c echo.Context, r Request) {
	c.Echo().Logger.Debug(r)
}
