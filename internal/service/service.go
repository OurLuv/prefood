package service

import "github.com/OurLuv/prefood/internal/storage/postgres"

type Service struct {
	CategoryService
	FoodService
	UserService
	RestaruantService
	OrderService
}

func NewService(repo postgres.Repository) *Service {
	return &Service{
		NewCategoryServiceImpl(repo.CategoryStorage),
		NewFoodServiceImpl(repo.FoodStorage),
		NewUserServiceImpl(repo.UserStorage),
		NewRestaruantServiceImpl(repo.RestaurantStorage),
		NewOrderServiceImpl(repo.OrderStorage),
	}
}
