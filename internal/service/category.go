package service

import (
	"context"

	"github.com/davisande/grpc-study-project/internal/database"
	"github.com/davisande/grpc-study-project/internal/pb"
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

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryRespose := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryRespose, nil
}
