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
		dsn := "host=gomud-db-postgres.gomud.svc.cluster.local user=gomud password=admin123 dbname=gomud port=5432 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Println("Error creating connection pool: " + err.Error())
			log.Println("Waiting for 5 seconds before retrying")
		} else {
			return
		}

		select {
		case <-ctx.Done():
		case <-time.After(5 * time.Second):
		}
	}
}
