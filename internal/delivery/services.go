package delivery

import (
	serviceCategory "category-service/internal/domain/category/services"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	CategoryService serviceCategory.CategoryService
}

type Deps struct {
	Repository *Repositories
	Redis      redis.Client
	//GRPC       *GRPC
}

func NewService(deps Deps) *Service {
	return &Service{
		CategoryService: serviceCategory.NewServiceCategory(deps.Repository.CategoryRepository, deps.Redis),
	}
}
