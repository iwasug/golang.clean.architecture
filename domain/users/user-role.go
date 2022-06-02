package users

import "time"

type UserRole struct {
	Id        string    `gorm:"column:RoleId;type:uuid;primary_key"`
	Name      string    `gorm:"column:Name"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	CreatedBy string    `gorm:"column:CreatedBy"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
	UpdatedBy string    `gorm:"column:UpdatedBy"`
	IsActive  bool      `gorm:"column:IsActive"`
}

var (
	UserRoleSa = &UserRole{Name: "Super Admin"}
)
