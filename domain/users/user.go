package users

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.clean.architecture/application/users/models"
	"golang.clean.architecture/domain/common"
)

type User struct {
	Id        string      `gorm:"column:Id;type:uuid;primary_key"`
	Fullname  string      `gorm:"column:Fullname"`
	UserName  string      `gorm:"column:Username;index"`
	Password  string      `gorm:"column:Password"`
	Roles     []*UserRole `gorm:"foreignKey:RoleId"`
	CreatedAt time.Time   `gorm:"column:CreatedAt"`
	CreatedBy string      `gorm:"column:CreatedBy"`
	UpdatedAt time.Time   `gorm:"column:UpdatedAt"`
	UpdatedBy string      `gorm:"column:UpdatedBy"`
	IsActive  bool        `gorm:"column:IsActive"`
}

func NewUser(model *models.NewUserModel) *User {

	var user *User

	if common.IsNullOrEmpty(model.Username) {
		panic(common.IsNullOrEmptyError("username"))
	}

	if common.IsNullOrEmpty(model.Fullname) {
		panic(common.IsNullOrEmptyError("fullname"))
	}

	hashedPwd, err := HashAndSalt([]byte(model.Password))

	if err != nil {
		log.Fatalln(err)
	}

	user = &User{
		Id:        uuid.New().String(),
		UserName:  model.Username,
		Password:  hashedPwd,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

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
