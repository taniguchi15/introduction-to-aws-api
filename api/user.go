package api

import (
	"net/http"

	"github.com/labstack/echo"
)

/**
 * [GET] /
 */
func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

/**
 * [POST] /user
 */
func postUser(c echo.Context) error {

	user := new(User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log(c, user)
	return c.NoContent(http.StatusOK)
}
