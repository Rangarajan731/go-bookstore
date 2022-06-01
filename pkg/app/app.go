package app

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
