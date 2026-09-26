package models

type ModelHasRole struct{
	RoleID			int64				`json:"role_id"`
	ModelID			int64				`json:"model_id"`
	ModelType 		string				`json:"model_type"`

	Role 			Role				`json:"role"`
}