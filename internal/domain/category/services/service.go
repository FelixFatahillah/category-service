package services

import (
	"category-service/internal/domain/category/dtos"
	"category-service/internal/domain/category/models"
	"category-service/internal/domain/category/repositories"
	"category-service/pkg/helper"
	"context"
	"github.com/redis/go-redis/v9"
)

type serviceCategory struct {
	Repository repositories.CategoryRepository
	Redis      redis.Client
}

type CategoryService interface {
	Create(ctx context.Context, input dtos.CreateCategoryDto) (*dtos.CreateCategoryResponseDto, error)
	GetAll(ctx context.Context, filter dtos.CategoryFilter) ([]models.Category, *helper.PaginationMeta, error)
	FindById(ctx context.Context, id string) (*models.Category, error)
	Update(ctx context.Context, input dtos.UpdateCategoryDto) error
	Delete(ctx context.Context, id string) error
}
