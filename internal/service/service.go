package service

import "github.com/OurLuv/prefood/internal/storage/postgres"

type Service struct {
	CategoryService
}

func NewService(repo postgres.Repository) *Service {
	return &Service{
		NewCategoryServiceImpl(repo.CategoryStorage),
	}
}