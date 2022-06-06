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
	Roles     []*UserRole `Gorm:"many2many: user_roles;"`
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

	hashedPwd, err := common.HashAndSalt([]byte(model.Password))

	if err != nil {
		log.Fatalln(err)
	}

	user = &User{
		Id:        uuid.New().String(),
		UserName:  model.Username,
		Password:  hashedPwd,
		Fullname:  model.Fullname,
		Roles:     NewRole(model.Roles),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user
}

func (u *User) ChangePassword(oldPassword, newPassword string) {

	if !common.ComparePasswords(oldPassword, []byte(newPassword)) {
		panic("")
	}

	hashedPwd, err := common.HashAndSalt([]byte(newPassword))

	if err != nil {
		log.Fatalln(err)
	}

	u.Password = hashedPwd
}
