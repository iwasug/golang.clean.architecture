package models

type NewUserModel struct {
	Username string       `json:"username"`
	Fullname string       `json:"fullname"`
	Password string       `json:"password"`
	Roles    []*RoleModel `json:"roles"`
}
