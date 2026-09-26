package models

import (
	"github.com/KhaledMo94/quick-discount/internal/types"
	"gorm.io/gorm"
)

type BannerableType string

const (
    BannerableServiceProvider  BannerableType = "App\\Models\\ServiceProvider"
    BannerableCategory BannerableType = "App\\Models\\Category"
    BannerableMainCategory BannerableType = "App\\Models\\MainCategory"
	BannerableService  BannerableType = "App\\Models\\Service"
)

var BannerableTypes = []BannerableType{
	BannerableServiceProvider,
	BannerableCategory,
	BannerableMainCategory,
	BannerableService,
}

type Banner struct{
	BaseModel

	Title              *string         `json:"title" db:"title"`
	Image              types.Image     `json:"image" db:"image" gorm:"embedded"`
	IsGoldenMembership bool            `json:"is_golden_membership" db:"is_golden_membership" gorm:"not null;default:false"`

	BannerableID   *int64          `json:"bannerable_id" db:"bannerable_id" gorm:"index"`
	BannerableType *BannerableType `json:"bannerable_type" db:"bannerable_type" gorm:"type:varchar(255);index"`

	Service         *Service         `json:"service" db:"-" gorm:"-"`
	ServiceProvider *ServiceProvider `json:"service_provider" db:"-" gorm:"-"`
	Category        *Category        `json:"category" db:"-" gorm:"-"`
	MainCategory    *MainCategory    `json:"main_category" db:"-" gorm:"-"`
}

type BannerSqlDTO struct {
	ID                 int64          `json:"id"`
    Title              *string        `json:"title"`
    Image              types.Image    `json:"image" gorm:"embedded"`
    IsGoldenMembership bool           `json:"is_golden_membership"`
    BannerableID       *int64         `json:"bannerable_id"`
    BannerableType     *BannerableType `json:"bannerable_type"`

    // Unified target details from the join
    TargetName string `json:"target_name"`
}

// by in-memory mapping
func LoadMorph(db *gorm.DB) ([]Banner , error){
	var banners []Banner
	if err:=db.Find(&banners).Error; err!=nil{
		return nil , err
	}

	if len(banners) > 0 {
		serviceIDs := []int64{}
		categoryIDs := []int64{}
		serviceProviderIDs := []int64{}
		mainCategoryIDs := []int64{}
	
		for _, c := range banners {
			if c.BannerableType == nil || c.BannerableID == nil {
				continue
			}
			switch *c.BannerableType {
			case BannerableCategory:
				categoryIDs = append(categoryIDs, *c.BannerableID)
			case BannerableServiceProvider:
				serviceProviderIDs = append(serviceProviderIDs, *c.BannerableID)
			case BannerableMainCategory:
				mainCategoryIDs = append(mainCategoryIDs, *c.BannerableID)
			case BannerableService:
				serviceIDs = append(serviceIDs, *c.BannerableID)
			}
		}
	
		var mainCategories   []MainCategory
		var categories       []Category
		var services         []Service
		var serviceProviders []ServiceProvider

		mainCategoriesMap := make(map[int64]MainCategory)
		categoriesMap := make(map[int64]Category)
		servicesMap := make(map[int64]Service)
		serviceProvidersMap := make(map[int64]ServiceProvider)

		if len(mainCategoryIDs) > 0 {
			db.Where("id IN ?", mainCategoryIDs).Find(&mainCategories)
			for _, m := range mainCategories {
				mainCategoriesMap[m.ID] = m
			}
		}
		if len(categoryIDs) > 0 {
			db.Where("id IN ?", categoryIDs).Find(&categories)
			for _, c := range categories {
				categoriesMap[c.ID] = c
			}
		}
		if len(serviceIDs) > 0 {
			db.Where("id IN ?", serviceIDs).Find(&services)
			for _, s := range services {
				servicesMap[s.ID] = s
			}
		}
		if len(serviceProviderIDs) > 0 {
			db.Where("id IN ?", serviceProviderIDs).Find(&serviceProviders)
			for _, sp := range serviceProviders {
				serviceProvidersMap[sp.ID] = sp
			}
		}

		for i := range banners {
			if banners[i].BannerableType == nil || banners[i].BannerableID == nil {
				continue
			}
			id := *banners[i].BannerableID
			for _, bt := range BannerableTypes {
				if *banners[i].BannerableType != bt {
					continue
				}
				switch bt {
				case BannerableMainCategory:
					if m, ok := mainCategoriesMap[id]; ok {
						banners[i].MainCategory = &m
					}
				case BannerableCategory:
					if c, ok := categoriesMap[id]; ok {
						banners[i].Category = &c
					}
				case BannerableService:
					if s, ok := servicesMap[id]; ok {
						banners[i].Service = &s
					}
				case BannerableServiceProvider:
					if sp, ok := serviceProvidersMap[id]; ok {
						banners[i].ServiceProvider = &sp
					}
				}
				break
			}
		}
	}

	return banners, nil
}

func IsGoldenMembershipScope(q *gorm.DB){
	q.Where("is_golden_membership = 1")
}

func IsNotGoldenMembershipScope(q *gorm.DB){
	q.Where("is_golden_membership = 0")
}

func LoadSqlMorph(q *gorm.DB) ([]BannerSqlDTO , error){
	var banners []BannerSqlDTO

	query := `
		SELECT 
		b.id ,
		b.title,
		b.is_golden_membership ,
		b.bannerable_type , 
		b.bannerable_id,
		b.image ,
		COALESCE(sp.name, c.name, mc.name, s.name, '') AS target_name
        FROM banners b
        LEFT JOIN service_providers sp 
            ON b.bannerable_type = ? AND b.bannerable_id = sp.id
        LEFT JOIN categories c 
            ON b.bannerable_type = ? AND b.bannerable_id = c.id
        LEFT JOIN main_categories mc 
            ON b.bannerable_type = ? AND b.bannerable_id = mc.id
        LEFT JOIN services s 
            ON b.bannerable_type = ? AND b.bannerable_id = s.id
	`

	err := 	q.Raw(query, 
		BannerableServiceProvider, 
        BannerableCategory, 
        BannerableMainCategory, 
        BannerableService).Scan(&banners).Error

	if err != nil {
		return nil , err
	}

	return banners , nil
} 

