package users

import (
	"context"

	"golang.clean.architecture/domain/users"
	"gorm.io/gorm"
)

const _tableName = "users"

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) users.IUserRepository {
	return &userRepository{db: db}
}

func (repository userRepository) FindOneById(ctx context.Context, id string) (*users.User, error) {
	var user *users.User
	repository.db.First(&user, id)
	err := repository.db.Table(_tableName).Where("'Id' = ?", id).First(&user)
	return user, err.Error
}

func (repository userRepository) FindOneByUsername(ctx context.Context, username string) (*users.User, error) {
	var user *users.User
	err := repository.db.Table(_tableName).Where("Username = ?", username).First(&user)
	return user, err.Error
}

func (repository userRepository) Add(ctx context.Context, user *users.User) error {
	err := repository.db.Table(_tableName).Create(&user)
	if err != nil {
		return err.Error
	}
	return nil
}

func (repository userRepository) Update(ctx context.Context, user *users.User) error {
	err := repository.db.Table(_tableName).Save(&user)
	if err != nil {
		return err.Error
	}
	return nil
}
