package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type ParticipationRepository interface {
	GetParticipation() (participation []models.Participation, err error)
	RegisterTurnament(participan *models.Participation) error
	CheckPartisipasion(participation *models.Participation) (err error)
}

type participationRepository struct {
	db *gorm.DB
}

func NewParticipationRepository(db *gorm.DB) *participationRepository {
	return &participationRepository{db}
}

func (p *participationRepository) GetParticipation() (participation []models.Participation, err error) {

	if err := config.DB.Preload("Player").Preload("Turnament").Find(&participation).Error; err != nil {
		return nil, err
	}

	return participation, nil
}

func (p *participationRepository) CheckPartisipasion(participation *models.Participation) (err error) {

	if err := config.DB.Where("player_id = ? AND turnament_id = ?", participation.PlayerID, participation.TurnamentID).First(&participation).Error; err != nil {
		return err
	}

	return nil
}

func (p *participationRepository) RegisterTurnament(participan *models.Participation) error {
	if err := config.DB.Save(&participan).Error; err != nil {
		return err
	}

	return nil
}
