package repository

import (
	"github.com/KhaledMo94/quick-discount/internal/db"
	"github.com/KhaledMo94/quick-discount/internal/models"
	"github.com/KhaledMo94/quick-discount/internal/types"
)

type MainCategoryRepository struct{
	db *db.GormDB
}

func NewMainCategoryRepository(db *db.GormDB) *MainCategoryRepository{
	return &MainCategoryRepository{db: db}
}

func (r *MainCategoryRepository) Create(mainCategory *models.MainCategory) error{
	return r.db.Create(mainCategory).Error
}

func(r *MainCategoryRepository) FindByID(id int64) (*models.MainCategory , error){
	var mainCategory models.MainCategory

	if err:= r.db.First(&mainCategory , id).Error; err!=nil{
		return nil , err
	}

	return &mainCategory , nil
}

func (r *MainCategoryRepository) FindBySlug(slug string) (*models.MainCategory, error) {
	var mainCategory models.MainCategory

	err := r.db.
		Scopes(types.ScopeActive).
		Where("(slug_ar = ? OR slug_en = ?)", slug, slug).
		First(&mainCategory).Error
	if err != nil {
		return nil, err
	}

	return &mainCategory, nil
}

func (r *MainCategoryRepository) FindAll() ([]models.MainCategory, error) {
	var mainCategories []models.MainCategory

	if err := r.db.Scopes(types.ScopeActive).Find(&mainCategories).Error; err != nil {
		return nil, err
	}

	return mainCategories, nil
}

func (r *MainCategoryRepository) Update(id int64, mainCategory *models.MainCategory) error{
	return r.db.Model(&models.MainCategory{}).
		Where("id = ?",id).
		Updates(mainCategory).Error
}

func (r *MainCategoryRepository) Delete(id int64) error {
	return r.db.Delete(&models.MainCategory{}, id).Error
}