package routes

import (
	"PongPedia/controllers"
	m "PongPedia/middleware"
	"PongPedia/repository/database"
	"PongPedia/usecase"
	"PongPedia/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	userRepository := database.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase, userRepository)

	playerRepository := database.NewPlayerRespository(db)
	playerUsecase := usecase.NewPlayerUsecase(playerRepository, userRepository)
	playerController := controllers.NewPlayerController(playerUsecase, playerRepository)

	authRepository := database.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository, userRepository)
	authController := controllers.NewAuthController(authUsecase, authRepository, userUsecase)
	participationRepository := database.NewParticipationRepository(db)

	turnamentRepository := database.NewTurnamentRepository(db)
	turnamentUsecase := usecase.NewTurnamentUsecase(turnamentRepository, playerRepository, userRepository, participationRepository)
	turnamentController := controllers.NewTurnamentControllers(turnamentUsecase, turnamentRepository)

	matchRepository := database.NewMatchRepository(db)
	matchUsecase := usecase.NewMatchUsecase(matchRepository, participationRepository)
	matchController := controllers.NewMatchController(matchUsecase, matchRepository)

	adminUsecase := usecase.NewDashboardUsecase(userRepository, turnamentRepository, matchRepository, playerRepository)
	adminController := controllers.NewAdminControllers(adminUsecase, matchUsecase, turnamentUsecase)

	// Validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.POST("/register", authController.RegisterUserController)
	e.POST("/login", authController.LoginUserController)

	// Admin Routes
	a := e.Group("Dashboard/Admin", m.IsLoggedIn)
	a.GET("", adminController.DashboardAdminController)
	a.GET("/user", adminController.GetAllUserController)
	a.POST("/turnament", adminController.CreateTurnamentController)
	a.PUT("/turnament/:id", adminController.UpdateTurnamentController)
	a.POST("/match", adminController.CreateMatchController)
	a.PUT("/match/:id", adminController.UpdateMatchController)

	// User Routes
	pf := e.Group("/profile", m.IsLoggedIn)
	pf.GET("", userController.GetUserController)
	pf.PUT("", userController.UpdateUserController)
	pf.DELETE("", userController.DeleteUserController)

	// User Player Routes
	pp := e.Group("/profile/player", m.IsLoggedIn)
	pp.PUT("", playerController.UpdatePlayerController)

	// Turnament Routes
	tt := e.Group("/tournament")
	tt.GET("", turnamentController.GetTurnamentController)
	tt.GET("/:id", turnamentController.GetTurnamentDetailController)
	tt.POST("/register", turnamentController.RegisterTurnamentController, m.IsLoggedIn)

	// Match Routes
	mm := e.Group("/match")
	mm.GET("", matchController.GetMatchController)
	mm.GET("/:id", matchController.GetMatchByIdController)

}
