package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type FoodService interface {
	Create(c model.Food) (*model.Food, error)
	GetById(restaurantId uint, id uint) (*model.Food, error)
	GetAll(id uint) ([]model.Food, error)
	UpdateById(c model.Food) error
	DeleteById(id uint) error
}

type UserService interface {
	Create(u model.User) error
	GetById(id uint) (*model.User, error)
	Login(email string, password string) (*model.User, error)
	UpdateById(id uint, c model.User) error
	DeleteById(id uint) error
	CheckForEmail(email string) error
}

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
