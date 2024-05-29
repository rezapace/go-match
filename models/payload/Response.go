package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProfileResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Player   PlayerResponse
}

type PlayerResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	UserID    int    `json:"-"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type TurnamentDetailResponse struct {
	Name          string `json:"name" form:"name"`
	StartDate     string `json:"start_date" form:"start_date"`
	EndDate       string `json:"end_date" form:"end_date"`
	Location      string `json:"location" form:"location"`
	Paticipations []ParticipationResponse
}
type TurnamentResponse struct {
	ID        uint   `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	Location  string `json:"location" form:"location"`
	Champions string `json:"champion" form:"champion"`
	Slot      int    `json:"slot" form:"slot"`
}

type UpdateMatchResponse struct {
	MatchName      string `json:"match_name"`
	MatchDate      string `json:"match_date"`
	Player_1       int    `json:"player_1"`
	Player_2       int    `json:"player_2"`
	Player_1_Score int    `json:"player_1_score"`
	Player_2_Score int    `json:"player_2_score"`
	TurnamentID    int    `json:"turnament_id"`
}

type ParticipationResponse struct {
	PlayerID int `json:"player_id"`
	Player   PlayerResponse
}

type DashboardAdminResponse struct {
	TotalUser      int64 `json:"total_user"`
	TotalPlayer    int64 `json:"total_player"`
	TotalTurnament int64 `json:"total_turnament"`
	TotalMatch     int64 `json:"total_match"`
}
