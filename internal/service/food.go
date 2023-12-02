package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

type FoodServiceImpl struct {
	repo postgres.FoodStorage
}

func (fs *FoodServiceImpl) Create(f model.Food) (*model.Food, error) {
	return fs.repo.Create(f)
}
func (fs *FoodServiceImpl) GetById(restaurantId uint, id uint) (*model.Food, error) {
	return fs.repo.GetById(restaurantId, id)
}
func (fs *FoodServiceImpl) GetAll(id uint) ([]model.Food, error) {
	return fs.repo.GetAll(id)
}
func (fs *FoodServiceImpl) UpdateById(c model.Food) error {
	return fs.repo.UpdateById(c)
}
func (fs *FoodServiceImpl) DeleteById(id uint) error {
	return nil
}

func NewFoodServiceImpl(repo postgres.FoodStorage) *FoodServiceImpl {
	return &FoodServiceImpl{
		repo: repo,
	}
}
