package models

import "github.com/KhaledMo94/quick-discount/internal/types"

type BaseModel struct {
	ID                               int64         			`json:"id" gorm:"primaryKey"`
	types.Timestamp                  `gorm:"embedded"`
}

type CommonNameDescImage struct {
	Name        string      `json:"name" db:"name"`
	Description *string     `json:"description" db:"description"`
	Image       types.Image `json:"image" db:"image" gorm:"embedded"`
}