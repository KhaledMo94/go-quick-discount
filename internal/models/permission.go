package models


type Permission struct {
	Name			string				`json:"name" gorm:"uniqueIndex:permissions_name_guard_name_unique"`
	GuardName		string				`json:"guard_name" gorm:"uniqueIndex:permissions_name_guard_name_unique"`

	Roles			[]Role				`json:"roles" gorm:"many2many:role_has_permissions;"`
	ModelHasPermissions []ModelHasPermission `json:"model_has_permission" gorm:"foreignKey:PermissionID"`

	BaseModel
}