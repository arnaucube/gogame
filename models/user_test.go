package models

import (
	"testing"

	"github.com/arnaucube/gogame/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := database.New("127.0.0.1:27017", "gogametests")
	assert.Nil(t, err)

	user, err := NewUser(db, "user00", "password00", "user00@email.com")
	assert.Nil(t, err)
	assert.Equal(t, user.Name, "user00")
}
