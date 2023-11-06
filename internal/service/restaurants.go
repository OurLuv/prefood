package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

type RestaruantService interface {
	GetAll(client_id uint) ([]model.Restaurant, error)
	Create(r model.Restaurant) error
	GetById(id uint, client_id uint) (*model.Restaurant, error)
}

type RestaruantServiceImpl struct {
	repo postgres.RestaurantStorage
}

func (rsi *RestaruantServiceImpl) GetAll(client_id uint) ([]model.Restaurant, error) {
	return rsi.repo.GetAll(client_id)
}
func (rsi *RestaruantServiceImpl) Create(r model.Restaurant) error {
	return nil
}
func (rsi *RestaruantServiceImpl) GetById(id uint, client_id uint) (*model.Restaurant, error) {
	return nil, nil
}

func NewRestaruantServiceImpl(repo postgres.RestaurantStorage) *RestaruantServiceImpl {
	return &RestaruantServiceImpl{
		repo: repo,
	}
}
