package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"

	"github.com/labstack/echo"
)

type TurnamentUsecase interface {
	GetTurnament() ([]payload.TurnamentResponse, error)
	GetTurnamentById(id int) (turnament *models.Turnament, err error)
	CreateTurnament(req *payload.TurnamentRequest) (res payload.TurnamentRequest, err error)
	UpdateTurnament(id int, req *payload.UpdateTurnamentRequest) (res payload.TurnamentResponse, err error)
	RegisterTurnament(id int, req *payload.RegisterTurnamentRequest) error
}

type turnamentUsecase struct {
	turnamentRepository database.TurnamentRepository
	playerRepository    database.PlayerRespository
	userReposistory     database.UserRepository
	particapationRepo   database.ParticipationRepository
}

func NewTurnamentUsecase(
	turnamentRepository database.TurnamentRepository,
	playerRepository database.PlayerRespository,
	userReposistory database.UserRepository,
	participationRepo database.ParticipationRepository,
) TurnamentUsecase {
	return &turnamentUsecase{
		turnamentRepository,
		playerRepository,
		userReposistory,
		participationRepo,
	}
}

func (t *turnamentUsecase) GetTurnament() ([]payload.TurnamentResponse, error) {
	turnament, err := t.turnamentRepository.GetTurnament()
	if err != nil {
		return nil, err
	}

	res := []payload.TurnamentResponse{}
	for _, v := range turnament {
		res = append(res, payload.TurnamentResponse{
			ID:        v.ID,
			Name:      v.Name,
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			Location:  v.Location,
			Champions: v.Champion,
			Slot:      v.Slot,
		})
	}
	return res, nil
}

func (t *turnamentUsecase) GetTurnamentById(id int) (turnament *models.Turnament, err error) {
	// Check Turnament ID
	turnament, err = t.turnamentRepository.GetTurnamentById(id)
	if err != nil {
		echo.NewHTTPError(400, err.Error())
		return
	}

	return turnament, nil
}

func (t *turnamentUsecase) CreateTurnament(req *payload.TurnamentRequest) (res payload.TurnamentRequest, err error) {
	turnamentReq := &models.Turnament{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Location:  req.Location,
		Champion:  req.Champion,
		Slot:      req.Slot,
	}

	err = t.turnamentRepository.CreateTurnament(turnamentReq)
	if err != nil {
		echo.NewHTTPError(400, err.Error())
		return
	}

	res = payload.TurnamentRequest{
		Name:      turnamentReq.Name,
		StartDate: turnamentReq.StartDate,
		EndDate:   turnamentReq.EndDate,
		Location:  turnamentReq.Location,
		Champion:  turnamentReq.Champion,
		Slot:      turnamentReq.Slot,
	}

	return
}

func (t *turnamentUsecase) UpdateTurnament(id int, req *payload.UpdateTurnamentRequest) (res payload.TurnamentResponse, err error) {
	_, err = t.turnamentRepository.GetTurnamentById(id)
	if err != nil {
		return res, errors.New("Turnament not found")
	}

	turnamentReq := &models.Turnament{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Location:  req.Location,
		Champion:  req.Champion,
		Slot:      req.Slot,
	}

	err = t.turnamentRepository.UpdateTurnamenById(id, turnamentReq)
	if err != nil {
		echo.NewHTTPError(400, err.Error())
		return
	}

	res = payload.TurnamentResponse{
		ID:        turnamentReq.ID,
		Name:      turnamentReq.Name,
		StartDate: turnamentReq.StartDate,
		EndDate:   turnamentReq.EndDate,
		Location:  turnamentReq.Location,
		Champions: turnamentReq.Champion,
		Slot:      turnamentReq.Slot,
	}

	return
}

func (t *turnamentUsecase) RegisterTurnament(id int, req *payload.RegisterTurnamentRequest) error {

	player, err := t.playerRepository.GetPlayerId(id)
	if err != nil {
		return echo.NewHTTPError(400, "fill your player profile first")
	}

	regisReq := &models.Participation{
		PlayerID:    int(player.ID),
		TurnamentID: req.TurnamentID,
	}

	// Check slot availability
	turnament, err := t.turnamentRepository.GetTurnamentById(regisReq.TurnamentID)
	if err != nil {
		return echo.NewHTTPError(400, "Turnament not found")
	}

	if turnament.Slot == 0 {
		return echo.NewHTTPError(400, "Turnament slot is full")
	}

	// Check if user already registered
	err = t.particapationRepo.CheckPartisipasion(regisReq)
	if err == nil {
		return echo.NewHTTPError(400, "Player already registered")
	}

	// Register turnament
	err = t.particapationRepo.RegisterTurnament(regisReq)
	if err != nil {
		return err
	}

	// Update slot
	turnament.Slot -= 1

	err = t.turnamentRepository.UpdateTurnament(turnament)
	if err != nil {
		return err
	}

	return nil
}
