package api

import (
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.GET("/", getRoot)
	e.POST("/user", postUser)
}