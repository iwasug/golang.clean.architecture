package users

import (
	"golang.clean.architecture/application/users/models"
)

type UserRole struct {
	RoleId string `gorm:"column:RoleId;type:uuid;primary_key"`
	UserId string `gorm:"column:UserId;type:uuid;primary_key"`
	Role   Role   `gorm:"foreignKey:RoleId"`
	User   User   `gorm:"foreignKey:UserId"`
}

func NewRole(model []*models.RoleModel) []*UserRole {

	var roles []*UserRole

	// for _, role := range model {
	// 	roles = append(roles, &UserRole{
	// 		Id:        role.Id,
	// 		Name:      role.Name,
	// 		CreatedAt: time.Now(),
	// 		UpdatedAt: time.Now(),
	// 	})
	// }

	return roles
}
