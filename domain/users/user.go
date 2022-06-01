package users

import (
	"log"

	"github.com/google/uuid"
	"golang.clean.architecture/domain/common"
)

type User struct {
	Id        string      `json:"id" gorm:"column:Id;primary_key"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	UserName  string      `json:"username" gorm:"column:Username;index"`
	Password  string      `json:"password"`
	Roles     []*UserRole `json:"roles"`
}

func NewUser(firstName, lastName, username, password string) *User {

	var user *User

	if common.IsNullOrEmpty(username) {
		panic(common.IsNullOrEmptyError("username"))
	}

	hashedPwd, err := HashAndSalt([]byte(password))

	if err != nil {
		log.Fatalln(err)
	}

	user = &User{
		Id:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		UserName:  username,
		Password:  hashedPwd,
	}

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

	if !ComparePasswords(oldPassword, []byte(newPassword)) {
		panic("")
	}

	hashedPwd, err := HashAndSalt([]byte(newPassword))

	if err != nil {
		log.Fatalln(err)
	}

	u.Password = hashedPwd
}
