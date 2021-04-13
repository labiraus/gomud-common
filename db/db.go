package db

import (
	"context"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//Setup connects to database instance
func Setup(ctx context.Context) {
	go func() {
		<-ctx.Done()
	}()

	var err error
	dsn := "host=gomud-db-postgresql.gomud.svc.cluster.local user=gomud password=admin123 dbname=gomud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
}
