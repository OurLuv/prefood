package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

type CategoryService interface {
	Create(c model.Сategory) error
	GetById(id uint) (*model.Сategory, error)
	GetAll() ([]model.Сategory, error)
	UpdateCategoryById(id uint, c model.Сategory) error
	DeleteCategoryById(id uint) error
}

type CategoryServiceImpl struct{
	repo postgres.CategoryStorage
}

//* Create
func (cr *CategoryServiceImpl) Create(c model.Сategory) error{
	return cr.repo.Create(c)
}

//* Get category by id
func (cs *CategoryServiceImpl) GetById(id uint) (*model.Сategory, error){
	return cs.repo.GetById(id)
}

//* get all categories
func (cs *CategoryServiceImpl) GetAll() ([]model.Сategory, error){
	return cs.repo.GetAll()
}

//* update category
func (cs *CategoryServiceImpl) UpdateCategoryById(id uint, c model.Сategory) error{
	return nil
}

//* delete category by id
func (cs *CategoryServiceImpl) DeleteCategoryById(id uint) error{
	return cs.repo.DeleteCategoryById(id)
}



func NewCategoryServiceImpl (repo postgres.CategoryStorage) *CategoryServiceImpl{
	return &CategoryServiceImpl{repo: repo}
}