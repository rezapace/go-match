package controllers

import (
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"

	"github.com/labstack/echo"
)

type AuthController interface {
	LoginUserController(c echo.Context) error
	RegisterUserController(c echo.Context) error
}

type authControler struct {
	authUsecase    usecase.AuthUsecase
	authRepository database.AuthRepository
	userUsecase    usecase.UserUsecase
}

func NewAuthController(
	authUsecase usecase.AuthUsecase,
	authRepository database.AuthRepository,
	userUsecase usecase.UserUsecase,
) *authControler {
	return &authControler{
		authUsecase,
		authRepository,
		userUsecase,
	}
}

func (a *authControler) LoginUserController(c echo.Context) error {
	req := payload.LoginRequest{}

	c.Bind(&req)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := a.authUsecase.LoginUser(&req)

	if err != nil {
		return echo.NewHTTPError(400, "Invalid Email or Password")
	}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Data:    res,
	})
}

func (a *authControler) RegisterUserController(c echo.Context) error {
	req := payload.RegisterRequest{}

	c.Bind(&req)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty or Password must be 6 character")
	}

	err := a.userUsecase.CreateUser(&req)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Success Register",
	})

}
