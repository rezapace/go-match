package usecase

import (
	"PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error)
}

type authUsecase struct {
	authRepository database.AuthRepository
	userRepository database.UserRepository
}

func NewAuthUsecase(
	authRepository database.AuthRepository,
	userRepository database.UserRepository,
) *authUsecase {
	return &authUsecase{authRepository, userRepository}
}

// Logic for login user
func (a *authUsecase) LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error) {

	user, err := a.userRepository.GetuserByEmail(req.Email)
	if err != nil {
		echo.NewHTTPError(400, "Email not registered")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		echo.NewHTTPError(400, err.Error())
		return
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		echo.NewHTTPError(400, "Failed to generate token")
		return
	}

	user.Token = token

	res = payload.LoginResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}
