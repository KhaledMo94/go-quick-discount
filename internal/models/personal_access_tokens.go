package models

import (
	"time"

	// "gorm.io/gorm"
)

type Tokenable string

const(
	TOKENABLE_TYPE_USER Tokenable = "App\\Models\\User"
)

var TokenableTypes  = []Tokenable{
	TOKENABLE_TYPE_USER,
}

type PersonalAccessToken struct {
	BaseModel

	TokenableID 	int64			`db:"tokenable_id" json:"-" gorm:"index:personal_access_tokens_tokenable_type_tokenable_id_index;"`
	TokenableType 	Tokenable		`db:"tokenable_type" json:"tokenable_type" gorm:"index:personal_access_tokens_tokenable_type_tokenable_id_index;"`
	Name  			string			`db:"-"`
	Token 			string			`db:"token" json:"-" gorm:"uniqueIndex:personal_access_tokens_token_unique;"`		
	LastUsedAt 		*time.Time		`db:"last_used_at"`
	ExpiresAt 		*time.Time		`db:"expires_at"`
}


