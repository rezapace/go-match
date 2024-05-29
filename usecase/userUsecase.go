package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetUserById(id int) (res payload.ProfileResponse, err error)
	UpdateUser(id int, req *payload.UpdateUserRequest) (res payload.UpdateUserRequest, err error)
	CreateUser(reqs *payload.RegisterRequest) error
	DeleteUser(id int, password string) error
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepository database.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

// Logic for get user with cookie
func (u *userUsecase) GetUserById(id int) (res payload.ProfileResponse, err error) {
	user, err := u.userRepository.GetUseById(id)

	if err != nil {
		echo.NewHTTPError(401, "This routes for user only")
		return
	}

	res = payload.ProfileResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Player: payload.PlayerResponse{
			ID:        user.Player.ID,
			Name:      user.Player.Name,
			Age:       user.Player.Age,
			BirthDate: user.Player.BirthDate,
			Gender:    user.Player.Gender,
		},
	}

	return res, nil
}

// Logic for update user
func (u *userUsecase) UpdateUser(id int, req *payload.UpdateUserRequest) (res payload.UpdateUserRequest, err error) {
	userReq := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	userReq.ID = uint(id)

	err = u.userRepository.UpdateUser(userReq)

	if err != nil {
		echo.NewHTTPError(400, "Failed to update user")
		return
	}

	res = payload.UpdateUserRequest{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	return
}

// Logic for create user
func (u *userUsecase) CreateUser(reqs *payload.RegisterRequest) error {

	userReq := &models.User{
		Username: reqs.Username,
		Email:    reqs.Email,
		Password: reqs.Password,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(400, "Failed to hash password")
	}

	userReq.Password = string(passwordHash)

	err = u.userRepository.CreateUser(userReq)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// Logic for Delete user
func (u *userUsecase) DeleteUser(id int, password string) error {

	user, err := u.userRepository.ReadToken(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return echo.NewHTTPError(400, err.Error())
	}

	err = u.userRepository.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}
