package types

import (
	"database/sql/driver"
	"errors"

	"gorm.io/gorm"
)

type Sex string
const (
	SexMale Sex = "m"
	SexFemale Sex = "f"
)

type NullSex struct{
	Sex Sex
	Valid bool
}

func (n *NullSex) Scan (value any) error {
	if value == nil {
		n.Sex , n.Valid = "" , false
		return nil
	}
	switch v:= value.(type){
	case string:
		n.Sex , n.Valid = Sex(v) , true
		
	case []byte :
		n.Sex , n.Valid = Sex(v) , true
		
	default :
		return errors.New("types: incompatible type for Sex")
	}
	return nil
}

func (n NullSex) Value() (driver.Value , error){
	if !n.Valid {
		return nil , nil
	}
	return string(n.Sex) , nil
}

func (n NullSex) Label() string {
	if !n.Valid {
		return ""
	}

	switch n.Sex {
	case SexFemale:
		return "Female"
	case SexMale:
		return "Male"
	default:
		return "Unknown"
	}
}

func (n NullSex) IsFemale() bool { return n.Valid && n.Sex == SexFemale }
func (n NullSex) IsMale() bool   { return n.Valid && n.Sex == SexMale }

func ScopeFemale(db *gorm.DB) *gorm.DB {
	return db.Where("sex = ?", SexFemale)
}

func ScopeMale(db *gorm.DB) *gorm.DB {
	return db.Where("sex = ?", SexMale)
}

func ScopeSexIsNull(db *gorm.DB) *gorm.DB {
	return db.Where("sex IS NULL")
}

type HasSex struct {
	Sex NullSex `json:"sex" db:"sex" gorm:"column:sex;type:char(1)"`
}
