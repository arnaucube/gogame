module github.com/arnaucube/gogame

go 1.12

require (
	github.com/appleboy/gin-jwt/v2 v2.6.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/color v1.7.0
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/urfave/cli v1.20.0
	golang.org/x/crypto v0.1.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

replace github.com/arnaucube/gogame => ./
