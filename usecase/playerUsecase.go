package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"

	"github.com/labstack/echo"
)

type PlayerUsecase interface {
	GetPlayer(id int) (*models.Player, error)
	UpdatePlayer(id int, req *payload.CreateUpdatePlayerRequest) error
}

type playerUsecase struct {
	playerRespository database.PlayerRespository
	userRepository    database.UserRepository
}

func NewPlayerUsecase(
	playerRespository database.PlayerRespository,
	userRepository database.UserRepository,
) *playerUsecase {
	return &playerUsecase{playerRespository, userRepository}
}

func (p *playerUsecase) GetPlayer(id int) (*models.Player, error) {

	user, err := p.userRepository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	player, err := p.playerRespository.GetPlayerId(int(user.ID))

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to get player")
	}

	return player, nil
}

func (p *playerUsecase) UpdatePlayer(id int, req *payload.CreateUpdatePlayerRequest) error {

	player, err := p.playerRespository.GetPlayerId(id)

	if err == nil {
		player.Name = req.Name
		player.Age = req.Age
		player.BirthDate = req.BirthDate
		player.Gender = req.Gender

		err = p.playerRespository.UpdatePlayer(player)
		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		userReq := &models.Player{
			Name:      req.Name,
			Age:       req.Age,
			BirthDate: req.BirthDate,
			Gender:    req.Gender,
			UserID:    id,
		}

		err = p.playerRespository.UpdatePlayer(userReq)
		if err != nil {
			return errors.New(err.Error())
		}
	}

	return nil
}
