package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.clean.architecture/domain/common"
	"golang.clean.architecture/domain/users/events"
)

type User struct {
	Id                primitive.ObjectID  `json:"id" bson:"_id"`
	FirstName         string              `json:"first_name"`
	LastName          string              `json:"last_name"`
	UserName          string              `json:"user_name"`
	EncryptedPassword *EncryptedPassword  `json:"encrypted_password"`
	Roles             []*UserRole         `json:"roles"`
	domainEvents      []common.IBaseEvent `json:"domain_events"`
}

func (u *User) ClearDomainEvents() {
	u.domainEvents = nil
}

func (u *User) GetDomainEvents() []common.IBaseEvent {
	return u.domainEvents
}

func (u *User) AddEvent(event common.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

func NewUser(firstName, lastName, username, password string) *User {

	var user *User

	if common.IsNullOrEmpty(username) {
		panic(common.IsNullOrEmptyError("username"))
	}

	user = &User{
		Id:                primitive.NewObjectID(),
		FirstName:         firstName,
		LastName:          lastName,
		UserName:          username,
		EncryptedPassword: NewEncryptedPassword(password),
	}

	user.AddEvent(&events.UserCreated{
		Id:        user.Id,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  username,
	})

	return user
}

func NewGuestUser() *User {

	user := NewUser("", "", "Guest", "12345")
	user.AddUserRole(UserRole_Guest)

	return user
}

func NewAdminUser(firstName, lastName, username, password string) *User {

	user := NewUser(firstName, lastName, username, password)
	user.AddUserRole(UserRole_Admin)

	return user
}

func (u *User) AddUserRole(role *UserRole) {

	if role == nil {
		panic(common.IsNullOrEmptyError("role"))
	}

	for _, roleItem := range u.Roles {
		if roleItem.Name == role.Name {
			panic(common.AlreadyExistRoleError(role.Name))
		}
	}

	u.Roles = append(u.Roles, role)
}

func (u *User) ChangePassword(oldPassword, newPassword string) {

	if !u.EncryptedPassword.VerifyPassword(oldPassword) {
		panic("")
	}

	u.EncryptedPassword = NewEncryptedPassword(newPassword)
}
