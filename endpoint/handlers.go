package endpoint

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
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
	Name     string `binding:"required"`
	Password string `binding:"required"`
	Email    string `binding:"required"`
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
