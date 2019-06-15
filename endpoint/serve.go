package endpoint

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/arnaucube/gogame/config"
	"github.com/arnaucube/gogame/constants"
	"github.com/arnaucube/gogame/database"
	"github.com/arnaucube/gogame/models"
	"github.com/arnaucube/gogame/services/gamesrv"
	"github.com/arnaucube/gogame/services/usersrv"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var serverConfig config.Config
var db *database.Db
var userservice *usersrv.Service
var gameservice *gamesrv.Service

func newApiService() *gin.Engine {
	api := gin.Default()
	api.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8080"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: constants.JWTIdKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					constants.JWTIdKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userid := bson.ObjectIdHex(claims[constants.JWTIdKey].(string))
			return &models.User{
				Id: userid,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginMsg LoginMsg
			if err := c.ShouldBind(&loginMsg); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			_, user, err := userservice.Login(loginMsg.Email, loginMsg.Password)
			if err != nil {
				fail(c, err, jwt.ErrFailedAuthentication.Error())
				return "", err
			}
			return user, nil

		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	api.GET("/info", handleInfo)
	api.POST("/register", handleRegister)
	// api.POST("/login", handleLogin)
	api.POST("/login", authMiddleware.LoginHandler)
	api.GET("/refresh_token", authMiddleware.RefreshHandler)

	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/", handleGetUser)
		api.GET("/resources", handleGetResources)
		api.GET("/planets", handleGetUserPlanets)
		api.GET("/planets/:planetid", handleGetPlanet)
		api.POST("/buildings", handlePostUpgradeBuilding)
	}
	return api
}

func Serve(cnfg config.Config, _db *database.Db, _userservice *usersrv.Service, _gameservice *gamesrv.Service) *gin.Engine {
	serverConfig = cnfg
	db = _db
	userservice = _userservice
	gameservice = _gameservice
	return newApiService()
}
