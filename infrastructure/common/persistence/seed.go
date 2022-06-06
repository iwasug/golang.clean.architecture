package persistence

import (
	"golang.clean.architecture/application/users/models"
	"golang.clean.architecture/domain/users"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {

	var id = "00000000-0000-0000-0000-000000000000"
	var roleSa = users.Role{Id: id, Name: "Super Admin"}
	var count int64
	db.Raw("SELECT COUNT(*) FROM roles Where \"RoleId\" = ?", id).Scan(&count)
	if count == 0 {
		db.Create(&roleSa)
	}

	db.Raw("SELECT COUNT(*) FROM users Where \"Id\" = ?", id).Scan(&count)
	if count == 0 {
		var roles = []*models.RoleModel{}
		roles = append(roles, &models.RoleModel{Id: roleSa.Id, Name: roleSa.Name})
		userSa := users.NewUser(&models.NewUserModel{
			Username: "sa@mail.com",
			Password: "Qwerty@1234",
			Roles:    roles,
			Fullname: "System Administrator",
		})
		userSa.Id = id
		db.Create(&userSa)
		for _, role := range roles {
			userRole := users.UserRole{UserId: userSa.Id, RoleId: role.Id}
			db.Create(userRole)
		}
	}

}
