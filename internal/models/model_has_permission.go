package models

type ModelHasPermission struct{
	PermissionID			int64 			`json:"permission_id" gorm:"primaryKey"`
	ModelType				string			`json:"model_type" gorm:"primaryKey;index:model_has_permissions_model_id_model_type_index"`
	ModelID					int64			`json:"model_id" gorm:"primaryKey;index:model_has_permissions_model_id_model_type_index"`

	Permission				Permission		`json:"permissions" gorm:"foreignKey:PermissionID"`
}