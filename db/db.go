package db

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//Setup connects to database instance
func Setup(ctx context.Context) {
	for {
		var err error
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  "host=service-db-postgres.gomud user=gomud password=admin123 dbname=gomud port=5432",
			PreferSimpleProtocol: true,
		}), &gorm.Config{})

		if err != nil {
			log.Println("Error creating connection pool: " + err.Error())
			log.Println("Waiting for 5 seconds before retrying")
		} else {
			return
		}

		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
			return
		}
	}
}
