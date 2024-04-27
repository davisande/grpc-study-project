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

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindById(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, err
}
