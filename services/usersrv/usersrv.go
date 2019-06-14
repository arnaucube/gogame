package usersrv

import (
	"errors"
	"time"

	"github.com/arnaucube/gogame/database"
	"github.com/arnaucube/gogame/models"
	"github.com/arnaucube/gogame/services/gamesrv"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	db      *database.Db
	gamesrv *gamesrv.Service
}

func New(db *database.Db, gameservice *gamesrv.Service) *Service {
	return &Service{
		db:      db,
		gamesrv: gameservice,
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (srv Service) Register(name, password, email string) (*models.User, error) {
	var userDb models.User
	err := srv.db.Users.Find(bson.M{"email": email}).One(&userDb)
	if err == nil {
		return nil, errors.New("user already exist")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := models.NewUser(srv.db, name, hashedPassword, email)
	if err != nil {
		return nil, err
	}

	_, _, err = srv.gamesrv.CreatePlanet(user.Id)

	return user, err
}

var signingKey = []byte("TODO") // TODO

func (srv Service) Login(email, password string) (*string, *models.User, error) {
	var userDb models.UserDb
	err := srv.db.Users.Find(bson.M{"email": email}).One(&userDb)
	if err != nil {
		return nil, nil, errors.New("user not exist")
	}
	if !checkPasswordHash(password, userDb.Password) {
		return nil, nil, errors.New("error with password")
	}
	user := models.UserDbToUser(srv.db, userDb)

	// create jwt
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["user"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return nil, nil, errors.New("error creating token")
	}

	return &tokenString, user, err
}

func (srv Service) GetUserById(userid string) (*models.User, error) {
	var userDb models.UserDb
	err := srv.db.Users.Find(bson.M{"_id": bson.ObjectIdHex(userid)}).One(&userDb)
	if err != nil {
		return nil, err
	}
	return models.UserDbToUser(srv.db, userDb), nil
}

// func (srv Service) GetUser(id bson.ObjectId) (*models.User, error) {
//         // update user stats
//         user := getUserFromDB
//         user.GetStats()
//
// }
//
// func (srv Service) GetUser(id bson.ObjectId) (*models.User, error) {
//         // update user stats
//
// }
