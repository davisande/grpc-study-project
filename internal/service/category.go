package service

import (
	"context"

	"github.com/davisande/grpc-study-project/internal/database"
	pb "github.com/davisande/grpc-study-project/proto"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categorydb database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categorydb,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryRespose := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryRespose,
	}, nil
}
