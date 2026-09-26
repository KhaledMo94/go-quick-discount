package models

import "github.com/KhaledMo94/quick-discount/internal/types"

type Category struct {
	BaseModel
	CommonNameDescImage

	BannerImage    types.Image `json:"banner_image" db:"banner_image" gorm:"embedded"`
	MainCategoryID *int64      `json:"main_category_id" db:"main_category_id"`
	types.HasStatus `gorm:"embedded"`

	MainCategory *MainCategory `json:"main_category" db:"-" gorm:"foreignKey:MainCategoryID"`
}
