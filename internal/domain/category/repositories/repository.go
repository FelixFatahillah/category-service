package repositories

import (
	"category-service/internal/domain/category/dtos"
	"category-service/internal/domain/category/models"
	"category-service/pkg/helper"
	"context"
	"gorm.io/gorm"
)

type repositoryCategory struct {
	db *gorm.DB
}

type CategoryRepository interface {
	Transaction(ctx context.Context, fn func(repo CategoryRepository) error) error
	Create(ctx context.Context, user models.Category) (*models.Category, error)
	GetAll(ctx context.Context, filter dtos.CategoryFilter) ([]models.Category, *helper.PaginationMeta, error)
	FindById(ctx context.Context, id string) (*models.Category, error)
	FindByCategory(ctx context.Context, category string) (*models.Category, error)
	Update(ctx context.Context, input *models.Category) error
	Delete(ctx context.Context, id string) error
}
