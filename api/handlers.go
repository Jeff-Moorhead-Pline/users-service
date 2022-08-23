package api

import (
	"net/http"

	"github.com/jeff-moorhead-pline/users-service/data"
	"github.com/labstack/echo/v4"
)

func AttachHandlers(e *echo.Echo) *echo.Echo {

	e.GET("/users/all", handleGetAllUsers)
	e.GET("/users/:username", handleGetUserByUsername)
	e.POST("/users/add", handleAddUser)
	e.PUT("/users/:username/update", handleUpdateUser)

	return e
}

func handleGetAllUsers(c echo.Context) error {

	u := data.GetUsersDataLayer()

	return c.JSON(http.StatusOK, u)
}

func handleGetUserByUsername(c echo.Context) error {

	u, err := data.FindUserDataLayer(c.Param("username"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func handleAddUser(c echo.Context) error {

	var u data.User
	if err := c.Bind(&u); err != nil {
		return err
	}

	if err := data.AddUserDataLayer(&u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, u)
}

func handleUpdateUser(c echo.Context) error {

	var u data.User
	if err := c.Bind(&u); err != nil {
		return err
	}

	username := c.Param("username")

	if err := data.UpdateUserDataLayer(&u, username); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, u)
}
