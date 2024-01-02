package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

type OrderService interface {
	Create(order model.Order) (*model.Order, error)
	GetAll(restaurant_id uint) ([]model.Order, error)
	GetById(restaurantId uint, order_Id uint) (*model.Order, error)
	Delete(id uint) error
	ChangeStatus(id uint, status string) (string, error)
}

type OrderServiceImpl struct {
	repo postgres.OrderStorage
}

func (os *OrderServiceImpl) Create(order model.Order) (*model.Order, error) {
	return os.repo.Create(order)
}

func (os *OrderServiceImpl) GetAll(restaurant_id uint) ([]model.Order, error) {
	return os.repo.GetAll(restaurant_id)
}
func (os *OrderServiceImpl) GetById(restaurantId uint, orderId uint) (*model.Order, error) {
	return os.repo.GetById(restaurantId, orderId)
}

func (os *OrderServiceImpl) Delete(id uint) error {
	return os.repo.Delete(id)
}
func (os *OrderServiceImpl) ChangeStatus(id uint, status string) (string, error) {
	return os.repo.ChangeStatus(id, status)
}

func NewOrderServiceImpl(repo postgres.OrderStorage) *OrderServiceImpl {
	return &OrderServiceImpl{
		repo: repo,
	}
}
