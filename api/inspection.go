package api

import (
	"net/http"

	"github.com/labstack/echo"
)

/**
 * [POST] /success
 */
func postSuccess(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

/**
 * [POST] /badrequest
 */
func postBadRequest(c echo.Context) error {
	return c.NoContent(http.StatusBadRequest)
}

/**
 * [POST] /servererror
 */
func postServerError(c echo.Context) error {
	return c.NoContent(http.StatusInternalServerError)
}
