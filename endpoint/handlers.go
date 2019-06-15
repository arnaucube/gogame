package endpoint

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/arnaucube/gogame/constants"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func fail(c *gin.Context, err error, msg string) {
	color.Red(msg + ": " + err.Error())
	c.JSON(400, gin.H{
		"error": msg,
	})
}
func handleMain(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "wellcome to gogame",
	})
}

func handleInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"info": "info",
	})
}

type RegisterMsg struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func handleRegister(c *gin.Context) {
	var registerMsg RegisterMsg
	c.BindJSON(&registerMsg)
	user, err := userservice.Register(registerMsg.Name, registerMsg.Password, registerMsg.Email)
	if err != nil {
		fail(c, err, "error on register")
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

type LoginMsg struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func handleLogin(c *gin.Context) {
	var loginMsg LoginMsg
	c.BindJSON(&loginMsg)
	token, user, err := userservice.Login(loginMsg.Email, loginMsg.Password)
	if err != nil {
		fail(c, err, "error on login")
		return
	}
	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}

func handleGetUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userid := bson.ObjectIdHex(claims[constants.JWTIdKey].(string))

	user, err := userservice.GetUserById(userid)
	if err != nil {
		fail(c, err, "error on getting user")
		return
	}
	// resources, err := user.GetResources()
	// if err != nil {
	//         fail(c, err, "error on getting user resources")
	//         return
	// }

	c.JSON(200, gin.H{
		"user": user,
		// "resources": resources,
	})
}

func handleGetResources(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userid := bson.ObjectIdHex(claims[constants.JWTIdKey].(string))

	user, err := userservice.GetUserById(userid)
	if err != nil {
		fail(c, err, "error on getting user")
		return
	}
	resources, err := user.GetResources()
	if err != nil {
		fail(c, err, "error on getting user resources")
		return
	}

	c.JSON(200, gin.H{
		"user":      user,
		"resources": resources,
	})
}

func handleGetUserPlanets(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userid := bson.ObjectIdHex(claims[constants.JWTIdKey].(string))

	planets, err := userservice.GetUserPlanetsById(userid)
	if err != nil {
		fail(c, err, "error on getting user planets")
		return
	}

	c.JSON(200, gin.H{
		"planets": planets,
	})
}

func handleGetPlanet(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userid := bson.ObjectIdHex(claims[constants.JWTIdKey].(string))
	planetid := c.Param("planetid")

	user, err := userservice.GetUserById(userid)
	if err != nil {
		fail(c, err, "error on getting user")
		return
	}

	planet, err := gameservice.GetBuildings(user, bson.ObjectIdHex(planetid))
	if err != nil {
		fail(c, err, "error upgrading building")
		return
	}

	c.JSON(200, gin.H{
		"planet": planet,
	})
}

type BuildMsg struct {
	PlanetId string
	Building string
}

func handlePostUpgradeBuilding(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userid := bson.ObjectIdHex(claims[constants.JWTIdKey].(string))
	var buildMsg BuildMsg
	err := c.BindJSON(&buildMsg)
	if err != nil {
		fail(c, err, "error parsing json")
		return
	}

	user, err := userservice.GetUserById(userid)
	if err != nil {
		fail(c, err, "error on getting user")
		return
	}

	planet, err := gameservice.UpgradeBuilding(user, bson.ObjectIdHex(buildMsg.PlanetId), buildMsg.Building)
	if err != nil {
		fail(c, err, "error upgrading building")
		return
	}

	c.JSON(200, gin.H{
		"planet": planet,
	})
}
