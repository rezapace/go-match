package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
)

type DashboardUsecase interface {
	DashboardAdmin() (res payload.DashboardAdminResponse)
	GetAllUser() (user []models.User, err error)
}

type dashboardUsecase struct {
	userRepo      database.UserRepository
	turnamemtRepo database.TurnamentRepository
	matchRepo     database.MatchRepository
	playerRepo    database.PlayerRespository
}

func NewDashboardUsecase(
	userRepository database.UserRepository,
	turnamemtRepository database.TurnamentRepository,
	matchRepository database.MatchRepository,
	playerRespository database.PlayerRespository,
) *dashboardUsecase {
	return &dashboardUsecase{
		userRepo:      userRepository,
		turnamemtRepo: turnamemtRepository,
		matchRepo:     matchRepository,
		playerRepo:    playerRespository,
	}
}

func (d *dashboardUsecase) DashboardAdmin() (res payload.DashboardAdminResponse) {

	totalUser := d.userRepo.CountUser()
	totalPlayer := d.playerRepo.CountPlayer()
	totalTurnament := d.turnamemtRepo.CountTurnament()
	totalMatch := d.matchRepo.CountMatch()

	res = payload.DashboardAdminResponse{
		TotalUser:      totalUser,
		TotalPlayer:    totalPlayer,
		TotalTurnament: totalTurnament,
		TotalMatch:     totalMatch,
	}

	return res
}

func (d *dashboardUsecase) GetAllUser() (user []models.User, err error) {

	users, err := d.userRepo.GetUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}
