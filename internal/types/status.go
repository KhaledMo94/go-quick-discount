package types

import(
	"gorm.io/gorm"
)
type Status string
const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

type HasStatus struct{
	Status Status `json:"status" db:"status" gorm:"column:status;type:varchar(20);default:active;index"`
}

func ScopeActive(q *gorm.DB) *gorm.DB {
	return q.Where("status = ?",StatusActive)
}

func ScopeInactive(q *gorm.DB) *gorm.DB {
	return q.Where("status = ?",StatusInactive)
}