package models

import "github.com/KhaledMo94/quick-discount/internal/types"

type MainCategory struct {
	BaseModel
	CommonNameDescImage
	types.HasStatus `gorm:"embedded"`

	SlugAr *string `json:"slug_ar" db:"slug_ar" gorm:"uniqueIndex:main_categories_slug_ar_unique"`
	SlugEn *string `json:"slug_en" db:"slug_en" gorm:"uniqueIndex:main_categories_slug_en_unique"`
}
