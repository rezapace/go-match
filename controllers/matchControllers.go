package controllers

import (
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type MatchController interface {
	GetMatchController(c echo.Context) error
	GetMatchByIdController(c echo.Context) error
}

type matchController struct {
	matchUsecase    usecase.MatchUsecase
	matchRepository database.MatchRepository
}

func NewMatchController(
	matchUsecase usecase.MatchUsecase,
	matchRepository database.MatchRepository) *matchController {
	return &matchController{matchUsecase, matchRepository}

}

func (m *matchController) GetMatchController(c echo.Context) error {
	match, err := m.matchUsecase.GetMatch()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Match",
		Data:    match,
	})
}

func (m *matchController) GetMatchByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	match, err := m.matchUsecase.GetMatchById(id)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Match",
		Data:    match,
	})
}
