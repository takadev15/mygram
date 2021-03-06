package databases

import (
	"fmt"
	"log"

	"github.com/takadev15/mygram-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	userName = "taka"
	dbName   = "hacktiv_mygram"
	dbPass   = "depok1001"
	dbPort   = "5432"
  dbHost   = "localhost"
	db       *gorm.DB
	err      error
)

func DBInit() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, userName, dbPass, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Databases Error", err.Error())
	}
	log.Printf("Databases Connected")
	db.Debug().AutoMigrate(models.Photo{}, models.User{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
  return db
}
