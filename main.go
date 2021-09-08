package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/taniguchi15/compar-nuxt-react-api/api"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.SetLevel(log.DEBUG)
	//Default {"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}
	e.Logger.SetHeader("${time_rfc3339} [${level}] (${short_file}:${line})")

	api.Route(e)

	e.Start(":3000")
}
