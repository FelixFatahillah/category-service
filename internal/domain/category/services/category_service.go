package services

import (
	"category-service/internal/domain/category/dtos"
	"category-service/internal/domain/category/models"
	"category-service/internal/domain/category/repositories"
	"category-service/pkg/exception"
	"category-service/pkg/helper"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func NewServiceCategory(repository repositories.CategoryRepository, redis redis.Client) *serviceCategory {
	return &serviceCategory{
		Repository: repository,
		Redis:      redis,
	}
}

func (service serviceCategory) Create(ctx context.Context, input dtos.CreateCategoryDto) (*dtos.CreateCategoryResponseDto, error) {
	categoryExist, _ := service.Repository.FindByCategory(ctx, input.Category)
	if categoryExist != nil {
		return nil, &exception.ErrWithCode{
			Code: http.StatusBadRequest,
			Err:  errors.New("category is already registered"),
		}
	}

	record, err := service.Repository.Create(ctx, models.Category{
		Category:    input.Category,
		Description: input.Description,
	})
	if err != nil {
		return nil, err
	}
	return &dtos.CreateCategoryResponseDto{
		Category:    record.Category,
		Description: record.Description,
	}, nil
}
func (service serviceCategory) GetAll(ctx context.Context, filter dtos.CategoryFilter) ([]models.Category, *helper.PaginationMeta, error) {
	return service.Repository.GetAll(ctx, filter)
}

func (service serviceCategory) FindById(ctx context.Context, id string) (*models.Category, error) {
	return service.Repository.FindById(ctx, id)
}

func (service serviceCategory) Update(ctx context.Context, input dtos.UpdateCategoryDto) error {
	err := service.Repository.Transaction(ctx, func(repo repositories.CategoryRepository) error {
		record, _ := service.Repository.FindById(ctx, input.ID)
		if record == nil {
			return &exception.ErrWithCode{
				Code: http.StatusNotFound,
				Err:  errors.New("category not found"),
			}
		}
		_ = service.Repository.Update(ctx,
			&models.Category{
				ID:          record.ID,
				Category:    input.Category,
				Description: input.Description,
			},
		)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (service serviceCategory) Delete(ctx context.Context, id string) error {
	category, _ := service.Repository.FindById(ctx, id)
	if category == nil {
		return &exception.ErrWithCode{
			Code: http.StatusNotFound,
			Err:  errors.New("category not found"),
		}
	}
	return service.Repository.Delete(ctx, id)
}
