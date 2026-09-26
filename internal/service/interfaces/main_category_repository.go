package interfaces

import(
	"github.com/KhaledMo94/quick-discount/internal/models"
)

type MainCategoryRepository interface {
	Create(mainCategory *models.MainCategory) (error)
	FindByID(id int64) (*models.MainCategory , error)
	FindBySlug(slug string) (*models.MainCategory , error)
	FindAll() ([]models.MainCategory , error)
	Update(id int64, mainCategory *models.MainCategory) error
	Delete(id int64) error
}