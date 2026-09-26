package types

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var allowedSortColumns = map[string]bool{
	"created_at": true,
	"updated_at": true,
	"name":       true,
}

type Timestamp struct{
	CreatedAt *time.Time `json:"created_at" db:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
}

func ScopeLatest(column string) func(db *gorm.DB) *gorm.DB{
	if !allowedSortColumns[column] {
		column = "created_at"
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{
			Column: clause.Column{Name: column},
		})
	}
}