package category

import (
	"category-service/internal/domain/category/services"
	"category-service/internal/infrastructure/pb"
	"context"
)

type CategoryHandlerGrpc struct {
	pb.UnimplementedCategoryServiceServer
	category services.CategoryService
}

func NewCategoryHandlerGrpcHandler(category services.CategoryService) *CategoryHandlerGrpc {
	categoryServer := CategoryHandlerGrpc{
		category: category,
	}
	return &categoryServer
}

func (grpc CategoryHandlerGrpc) GetCategoryById(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error) {
	record, err := grpc.category.FindById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetCategoryByIdResponse{
		Category:    record.Category,
		Description: record.Description,
	}, nil
}
