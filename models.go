package main

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbFile := "iam_auth.db"
	conn, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Error("error connecting to database", err)
		log.Fatal("failed to connect database")
	}
	DB = conn
}

func AutoMigrate() {
	DB.AutoMigrate(&User{})
}
