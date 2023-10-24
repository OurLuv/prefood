package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

type FoodService interface {
	Create(c model.Food) error
	GetById(id uint) (*model.Food, error)
	GetAll() ([]model.Food, error)
	UpdateById(id uint, c model.Food) error
	DeleteById(id uint) error
}

type FoodServiceImpl struct {
	repo postgres.FoodStorage
}

func (fs *FoodServiceImpl) Create(f model.Food) error {
	return fs.repo.Create(f)
}
func (fs *FoodServiceImpl) GetById(id uint) (*model.Food, error) {
	return nil, nil
}
func (fs *FoodServiceImpl) GetAll() ([]model.Food, error) {
	return fs.repo.GetAll()
}
func (fs *FoodServiceImpl) UpdateById(id uint, c model.Food) error {
	return nil
}
func (fs *FoodServiceImpl) DeleteById(id uint) error {
	return nil
}

func NewFoodServiceImpl(repo postgres.FoodStorage) *FoodServiceImpl {
	return &FoodServiceImpl{
		repo: repo,
	}
}
