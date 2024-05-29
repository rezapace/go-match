package models

type Participation struct {
	ID          int `json:"id" form:"id" gorm:"primary_key"`
	PlayerID    int `json:"player_id" form:"player_id"`
	TurnamentID int `json:"turnament_id" form:"turnament_id"`
	Player      Player
}
