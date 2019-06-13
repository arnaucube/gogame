package usersrv

import (
	"errors"
	"time"

	"github.com/arnaucube/gogame/database"
	"github.com/arnaucube/gogame/models"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	db *database.Db
}

func New(db *database.Db) *Service {
	return &Service{
		db,
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
	var user models.User
	err := srv.db.Users.Find(bson.M{"email": email}).One(&user)
	if err == nil {
		return nil, errors.New("user already exist")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser := models.User{
		Name:     name,
		Password: hashedPassword,
		Email:    email,
	}
	err = srv.db.Users.Insert(newUser)
	if err != nil {
		return nil, err
	}
	err = srv.db.Users.Find(bson.M{"email": email}).One(&user)
	user.Password = ""
	return &user, err
}

var signingKey = []byte("TODO") // TODO

func (srv Service) Login(email, password string) (*string, *models.User, error) {
	var user models.User
	err := srv.db.Users.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, nil, errors.New("user not exist")
	}
	if !checkPasswordHash(password, user.Password) {
		return nil, nil, errors.New("error with password")
	}
	user.Password = ""

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

	return &tokenString, &user, err
}
