package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"
	"fmt"

	"github.com/labstack/echo"
)

type UserController interface {
	GetUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
	CreatePlayerController(c echo.Context) error
	UpdatePlayerController(c echo.Context) error
}

type userController struct {
	userUsecase    usecase.UserUsecase
	userRepository database.UserRepository
}

func NewUserController(
	userUsecase usecase.UserUsecase,
	userRepository database.UserRepository,
) *userController {
	return &userController{userUsecase, userRepository}
}

func (u *userController) GetUserController(c echo.Context) error {

	id, err := m.IsUser(c)
	if err != nil {
		return echo.NewHTTPError(400, "This routes for user only")
	}

	res, err := u.userUsecase.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: fmt.Sprintf("Welcome %s", res.Username),
		Data:    res,
	})
}

func (u *userController) UpdateUserController(c echo.Context) error {
	req := payload.UpdateUserRequest{}

	id, err := m.IsUser(c)
	if err != nil {
		return echo.NewHTTPError(400, "this routes for user only")
	}

	c.Bind(&req)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := u.userUsecase.UpdateUser(id, &req)
	if err != nil {
		return echo.NewHTTPError(400, "failed to update user")
	}

	return c.JSON(200, payload.Response{
		Message: "Success update user",
		Data:    res,
	})
}

func (u *userController) DeleteUserController(c echo.Context) error {
	id, err := m.IsUser(c)
	if err != nil {
		return echo.NewHTTPError(400, "this routes for user only")
	}

	password := c.FormValue("Password")

	if err := u.userUsecase.DeleteUser(id, password); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, "Succes Delete User")
}
