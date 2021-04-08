package db

import (
	"context"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

//Setup connects to database instance
func Setup(ctx context.Context) {
	go func() {
		<-ctx.Done()
	}()

	var err error
	db, err = gorm.Open(sqlserver.Open("server=192.168.56.1;user id=sa;password=superStr0ng!Password;database=gomud"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
}
