package models

type Match struct {
	ID             int    `json:"id" form:"id" gorm:"primary_key"`
	MatchName      string `json:"match_name" form:"match_name"`
	MatchDate      string `json:"match_date" form:"match_date"`
	Player_1       int    `json:"player_1" form:"player_1"`
	Player_2       int    `json:"player_2" form:"player_2"`
	Player_1_Score int    `json:"player_1_score" form:"player_1_score"`
	Player_2_Score int    `json:"player_2_score" form:"player_2_score"`
	TurnamentID    int    `json:"turnament_id" form:"turnament_id"`
	Turnament      Turnament
}
