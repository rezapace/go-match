package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type PlayerRespository interface {
	GetPlayer() (player []models.Player, err error)
	UpdatePlayer(player *models.Player) error
	GetPlayerId(id int) (player *models.Player, err error)
	CountPlayer() (res int64)
}

type playerRespository struct {
	db *gorm.DB
}

func NewPlayerRespository(db *gorm.DB) *playerRespository {
	return &playerRespository{db}
}

func (p *playerRespository) GetPlayer() (player []models.Player, err error) {

	if err := config.DB.Preload("Participation").Find(&player).Error; err != nil {
		return nil, err
	}

	return player, nil
}

func (p *playerRespository) CountPlayer() (res int64) {
	res = 0
	player := []models.Player{}

	if err := config.DB.Model(&player).Count(&res).Error; err != nil {
		return 0
	}

	return res
}

func (p *playerRespository) GetPlayerId(id int) (player *models.Player, err error) {

	if err := config.DB.Where("user_id = ?", id).Preload("Participation").First(&player).Error; err != nil {
		return nil, err
	}

	return player, nil
}

func (p *playerRespository) UpdatePlayer(player *models.Player) error {

	if err := config.DB.Save(&player).Error; err != nil {
		return err
	}
	return nil
}
