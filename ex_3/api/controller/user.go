package controller

import (
	"ex_3/api/manager"
	"ex_3/api/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *Controller) getUser(ctx echo.Context) error {
	mng := manager.GetInstance()

	user, err := mng.GetUser(ctx.QueryParam("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (c *Controller) createUser(ctx echo.Context) error {
	mng := manager.GetInstance()

	user := new(model.User)

	if err := ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := mng.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, created)
}

func (c *Controller) updateUser(ctx echo.Context) error {
	mng := manager.GetInstance()

	user := new(model.User)

	if err := ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := mng.UpdateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, created)
}

func (c *Controller) deleteUser(ctx echo.Context) error {
	mng := manager.GetInstance()

	user := new(model.User)

	if err := ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	deleted, err := mng.DeleteUser(user)
	fmt.Println(deleted, err)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, deleted)
}
