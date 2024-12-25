package delivery

import (
	repositoryCategory "category-service/internal/domain/category/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	CategoryRepository repositoryCategory.CategoryRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		CategoryRepository: repositoryCategory.NewRepositoryCategory(db),
	}
}
