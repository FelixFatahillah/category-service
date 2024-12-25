package repositories

import (
	"category-service/internal/domain/category/dtos"
	"category-service/internal/domain/category/models"
	"category-service/pkg/helper"
	"category-service/pkg/logger"
	"context"
	"gorm.io/gorm"
)

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db}
}

func (repository repositoryCategory) beginTransaction() *gorm.DB { return repository.db.Begin() }

func (repository repositoryCategory) withTx(ctx context.Context, tx *gorm.DB) *repositoryCategory {
	repository.db = tx.WithContext(ctx)
	return &repository
}

func (repository repositoryCategory) Transaction(ctx context.Context, fn func(repo CategoryRepository) error) error {
	tx := repository.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	repo := repository.withTx(ctx, tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repository repositoryCategory) Create(ctx context.Context, user models.Category) (*models.Category, error) {
	if err := repository.db.WithContext(ctx).Create(&user).Error; err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func (repository repositoryCategory) GetAll(ctx context.Context, filter dtos.CategoryFilter) ([]models.Category, *helper.PaginationMeta, error) {
	records := make([]models.Category, 0)
	paginateMeta := helper.PaginationMeta{
		Page:  filter.Pagination.Page,
		Limit: filter.Pagination.Limit,
	}

	query := repository.db.WithContext(ctx).Model(models.Category{}).Scopes(helper.PaginateScope(&filter.Pagination))
	query.Count(&paginateMeta.Total).Order("created_date DESC").Find(&records)

	paginateMeta.TotalPage = helper.GetTotalPage(paginateMeta.Total, paginateMeta.Limit)

	return records, &paginateMeta, nil
}

func (repository repositoryCategory) FindById(ctx context.Context, id string) (*models.Category, error) {
	record := &models.Category{}
	err := repository.db.WithContext(ctx).Take(record, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (repository repositoryCategory) FindByCategory(ctx context.Context, category string) (*models.Category, error) {
	record := &models.Category{}
	err := repository.db.WithContext(ctx).Take(record, "category = ?", category).Error

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (repository repositoryCategory) Update(ctx context.Context, input *models.Category) error {
	if err := repository.db.
		WithContext(ctx).
		Updates(input).Error; err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (repository repositoryCategory) Delete(ctx context.Context, id string) error {
	err := repository.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&models.Category{}).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
