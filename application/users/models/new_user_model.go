package models

type NewUserModel struct {
	Id       string       `json:"id"`
	Username string       `json:"username"`
	Fullname string       `json:"fullname"`
	Password string       `json:"password"`
	Roles    []*RoleModel `json:"roles"`
}
