package persistence

import (
	"fmt"
	"log"

	"golang.clean.architecture/api/configs"
	domainUsers "golang.clean.architecture/domain/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionDb(config configs.Database) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domainUsers.User{})
	db.AutoMigrate(&domainUsers.UserRole{})

	return db
}
