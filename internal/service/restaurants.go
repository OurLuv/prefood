package service

import (
	"time"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
)

type RestaruantService interface {
	GetAll(client_id uint) ([]model.Restaurant, error)
	Create(r model.Restaurant) (uint, error)
	GetById(id uint, client_id uint) (*model.Restaurant, error)
	Update(r model.Restaurant, data ToUpdate) error
	Delete(id uint) error
	OpenClose(id uint) (*bool, error)
}

type ToUpdate struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type RestaruantServiceImpl struct {
	repo postgres.RestaurantStorage
}

func (rsi *RestaruantServiceImpl) GetAll(client_id uint) ([]model.Restaurant, error) {
	return rsi.repo.GetAll(client_id)
}
func (rsi *RestaruantServiceImpl) Create(r model.Restaurant) (uint, error) {
	r.CreatedAt = time.Now()
	return rsi.repo.Create(r)
}
func (rsi *RestaruantServiceImpl) GetById(id uint, client_id uint) (*model.Restaurant, error) {
	return rsi.repo.GetById(id, client_id)
}

func (rsi *RestaruantServiceImpl) Update(r model.Restaurant, d ToUpdate) error {
	r.Name = d.Name
	r.Phone = d.Phone
	r.Country = d.Country
	r.State = d.State
	r.City = d.City
	r.Street = d.Street
	return rsi.repo.Update(r)
}
func (rsi *RestaruantServiceImpl) Delete(id uint) error {
	return rsi.repo.Delete(id)
}
func (rsi *RestaruantServiceImpl) OpenClose(id uint) (*bool, error) {
	return rsi.repo.OpenClose(id)
}

func NewRestaruantServiceImpl(repo postgres.RestaurantStorage) *RestaruantServiceImpl {
	return &RestaruantServiceImpl{
		repo: repo,
	}
}
