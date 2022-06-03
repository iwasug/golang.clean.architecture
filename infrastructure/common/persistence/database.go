package persistence

import (
	"log"

	domainUsers "golang.clean.architecture/domain/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionDb(dbUrl string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domainUsers.User{})
	db.AutoMigrate(&domainUsers.UserRole{})

	return db
}
