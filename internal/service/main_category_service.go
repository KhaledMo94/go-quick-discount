package service

import (
	"github.com/KhaledMo94/quick-discount/internal/models"
	"github.com/KhaledMo94/quick-discount/internal/service/interfaces"
)

type MainCategoryService struct {
	r interfaces.MainCategoryRepository
}

func NewMainCategoryService(repo interfaces.MainCategoryRepository) *MainCategoryService {
	return &MainCategoryService{r: repo}
}

func (s *MainCategoryService) Create(mainCategory *models.MainCategory) error {
	return s.r.Create(mainCategory)
}

func (s *MainCategoryService) FindByID(id int64) (*models.MainCategory, error) {
	return s.r.FindByID(id)
}

func (s *MainCategoryService) FindBySlug(slug string) (*models.MainCategory, error) {
	return s.r.FindBySlug(slug)
}

func (s *MainCategoryService) FindAll() ([]models.MainCategory, error) {
	return s.r.FindAll()
}

func (s *MainCategoryService) Update(id int64, mainCategory *models.MainCategory) error {
	return s.r.Update(id, mainCategory)
}

func (s *MainCategoryService) Delete(id int64) error {
	return s.r.Delete(id)
}
