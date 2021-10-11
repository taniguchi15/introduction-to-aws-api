package api

import (
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.GET("/", getRoot)
	e.POST("/user", postUser)
	e.POST("/success", postSuccess)
	e.POST("/badrequest", postBadRequest)
	e.POST("/servererror", postServerError)
}
