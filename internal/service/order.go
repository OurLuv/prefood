package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

type OrderService interface {
	Create(order model.Order) error
	GetAll(restaurant_id uint) ([]model.Order, error)
	GetById(restaurantId uint, order_Id uint) (*model.Order, error)
}

type OrderServiceImpl struct {
	repo postgres.OrderStorage
}

func (os *OrderServiceImpl) Create(order model.Order) error {
	return os.repo.Create(order)
}

func (os *OrderServiceImpl) GetAll(restaurant_id uint) ([]model.Order, error) {
	return os.repo.GetAll(restaurant_id)
}
func (os *OrderServiceImpl) GetById(restaurantId uint, orderId uint) (*model.Order, error) {
	return os.repo.GetById(restaurantId, orderId)
}

func NewOrderServiceImpl(repo postgres.OrderStorage) *OrderServiceImpl {
	return &OrderServiceImpl{
		repo: repo,
	}
}
