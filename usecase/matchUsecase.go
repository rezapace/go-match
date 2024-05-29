package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type MatchUsecase interface {
	GetMatch() ([]models.Match, error)
	GetMatchById(id int) (*models.Match, error)
	CreateMatch(req *payload.CreateMatchRequest) error
	UpdateMatch(req *payload.UpdateMatchRequest, id int) (res payload.UpdateMatchResponse, err error)
}

type matchUsecase struct {
	matchRepository         database.MatchRepository
	participationRepository database.ParticipationRepository
}

func NewMatchUsecase(
	matchRepository database.MatchRepository,
	participationRepository database.ParticipationRepository,
) *matchUsecase {
	return &matchUsecase{matchRepository, participationRepository}
}

func (m *matchUsecase) GetMatch() ([]models.Match, error) {
	match, err := m.matchRepository.GetMatch()
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (m *matchUsecase) GetMatchById(id int) (*models.Match, error) {
	match, err := m.matchRepository.GetMatchById(id)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (m *matchUsecase) CreateMatch(req *payload.CreateMatchRequest) error {
	matchReq := &models.Match{
		MatchName:      req.MatchName,
		MatchDate:      req.MatchDate,
		Player_1:       req.Player_1,
		Player_2:       req.Player_2,
		Player_1_Score: req.Player_1_Score,
		Player_2_Score: req.Player_2_Score,
		TurnamentID:    req.TurnamentID,
	}

	_, err := m.matchRepository.Checkmatch(req.Player_1, req.Player_2)
	if err == nil {
		return echo.NewHTTPError(400, "Match already exist")
	}

	// check participation in turnament
	_, err = m.matchRepository.CheckPartisipant(req.TurnamentID, req.Player_1)
	if err != nil {
		return echo.NewHTTPError(400, "Player not participate in this turnament")
	}

	_, err = m.matchRepository.CheckPartisipant(req.TurnamentID, req.Player_2)
	if err != nil {
		return echo.NewHTTPError(400, "Player not participate in this turnament")
	}

	err = m.matchRepository.CreateMatch(matchReq)
	if err != nil {
		return err
	}

	return nil
}

func (m *matchUsecase) UpdateMatch(req *payload.UpdateMatchRequest, id int) (res payload.UpdateMatchResponse, err error) {
	match, err := m.matchRepository.GetMatchById(id)
	if err != nil {
		echo.NewHTTPError(400, "Match not found")
		return
	}

	match.Player_1_Score = req.Player_1_Score
	match.Player_2_Score = req.Player_2_Score

	err = m.matchRepository.UpdateMatch(id, match)
	if err != nil {
		return res, err
	}

	res = payload.UpdateMatchResponse{
		MatchName:      match.MatchName,
		MatchDate:      match.MatchDate,
		Player_1:       match.Player_1,
		Player_2:       match.Player_2,
		Player_1_Score: match.Player_1_Score,
		Player_2_Score: match.Player_2_Score,
		TurnamentID:    match.TurnamentID,
	}

	return res, nil
}
