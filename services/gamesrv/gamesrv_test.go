package gamesrv

import (
	"fmt"
	"testing"

	"github.com/arnaucube/gogame/database"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestCreatePlanet(t *testing.T) {
	db, err := database.New("127.0.0.1:27017", "gogametests")
	assert.Nil(t, err)
	srv := New(db)

	solarSystem, planet, err := srv.CreatePlanet(bson.ObjectIdHex("5d029a6ff18ba24f406168fe"))
	assert.Nil(t, err)
	fmt.Println(solarSystem)
	fmt.Println(planet)
}
