package models

type Role struct {
	Name      		string `json:"name" gorm:"uniqueIndex:roles_name_guard_name_unique"`
	GuardName 		string `json:"guard_name" gorm:"uniqueIndex:roles_name_guard_name_unique"`
	Permissions		[]Permission `json:"permissions" gorm:"many2many:role_has_permissions;"`

	BaseModel
}
