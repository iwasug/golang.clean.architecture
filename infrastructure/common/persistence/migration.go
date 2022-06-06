package persistence

import (
	"fmt"
	"log"

	"golang.clean.architecture/api/configs"
	domainUsers "golang.clean.architecture/domain/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AutoMigration(config configs.Database) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domainUsers.User{})
	db.AutoMigrate(&domainUsers.Role{})
	db.AutoMigrate(&domainUsers.UserRole{})

	//Seed Data
	Seed(db)

}
