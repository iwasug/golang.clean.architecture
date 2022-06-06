package users

import (
	"context"
)

type IUserRepository interface {
	FindOneById(ctx context.Context, id string) (*User, error)
	FindOneByUsername(ctx context.Context, username string) (*User, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}
