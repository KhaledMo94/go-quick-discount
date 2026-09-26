package models

import "github.com/KhaledMo94/quick-discount/internal/types"



type City struct{
	Name			string				`json:"name"`
	Description 	*string				`json:"description"`
	Image          types.Image     `json:"image" gorm:"embedded"`

	BaseModel

	//relations 
	Users			[]User			`json:"users"`
}