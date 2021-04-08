package db

import (
	"errors"

	"gorm.io/gorm"
)

//CreateUser creates a new user
func CreateUser(userName, password, characterName string) *User {
	user := &User{UserName: userName,
		CharacterName: characterName,
		PasswordHash:  hash(password),
	}
	db.Create(user)
	return user
}

//Login returns a user if it exists
func Login(userName, password string) *User {
	user := &User{}
	result := db.Where(&User{UserName: userName, PasswordHash: hash(password)}).Find(user)

	if user == nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return user
}

//UserNameExists checks to see if a user name exists
func UserNameExists(userName string) bool {
	var user User
	result := db.Where(&User{UserName: userName}).First(&user)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

//CharacterNameExists checks to see if a character name exists
func CharacterNameExists(characterName string) bool {
	var user User
	result := db.Where(&User{CharacterName: characterName}).First(&user)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
