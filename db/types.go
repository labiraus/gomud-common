package db

import (
	"github.com/labiraus/gomud-common.git/core"

	"gorm.io/gorm"
)

//Connection describes the connection between two rooms
type Connection struct {
	gorm.Model
	RoomID      uint
	NeighbourID uint
	Dir         core.Direction
	TwoWay      bool
}

//Room describes a room
type Room struct {
	gorm.Model
	Name        string
	Description string
	Connections []Connection
}

//User describes a logged on character
type User struct {
	gorm.Model
	UserName      string
	CharacterName string
	PasswordHash  string
}
