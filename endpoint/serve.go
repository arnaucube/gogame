package endpoint

import (
	"github.com/arnaucube/gogame/config"
	"github.com/arnaucube/gogame/database"
	"github.com/arnaucube/gogame/services/gamesrv"
	"github.com/arnaucube/gogame/services/usersrv"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var serverConfig config.Config
var db *database.Db
var userservice *usersrv.Service
var gameservice *gamesrv.Service

func newApiService() *gin.Engine {
	api := gin.Default()
	api.Use(cors.Default())
	api.GET("/", handleMain)
	api.GET("/info", handleInfo)
	api.POST("/register", handleRegister)
	api.POST("/login", handleLogin)

	// TODO add jwt checker
	api.GET("/resources/:userid", handleGetResources)
	api.GET("/planets/:userid", handleGetUserPlanets)
	api.POST("/buildings/:userid", handlePostUpgradeBuilding)
	return api
}

func Serve(cnfg config.Config, _db *database.Db, _userservice *usersrv.Service, _gameservice *gamesrv.Service) *gin.Engine {
	serverConfig = cnfg
	db = _db
	userservice = _userservice
	gameservice = _gameservice
	return newApiService()
}
