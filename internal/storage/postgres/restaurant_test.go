package postgres

import (
	"testing"
	"time"

	"github.com/OurLuv/prefood/internal/model"
)

func TestCreateRestaurant(t *testing.T) {
	repo := NewRestaurantRepository(pool)
	r := model.Restaurant{
		Name:      "New Subway",
		ClientId:  3,
		Phone:     "999-888-7777",
		Country:   "United Kingdom",
		State:     "England",
		City:      "London",
		Street:    "655 Oxford Street",
		Email:     "info@subway.co.uk",
		CreatedAt: time.Now(),
	}
	if _, err := repo.Create(r); err != nil {
		t.Error(err)
	}
}

func TestGetRestaurant(t *testing.T) {
	repo := NewRestaurantRepository(pool)

	r, err := repo.GetById(1, 3)
	if err != nil {
		t.Error(err)
	}
	_ = r
}

func TestUpdateRestaurant(t *testing.T) {
	repo := NewRestaurantRepository(pool)
	r := model.Restaurant{
		Id:        4,
		Name:      "Subway New Coast",
		ClientId:  3,
		Phone:     "999-888-7777",
		Country:   "United Kingdom",
		State:     "England",
		City:      "London",
		Street:    "665 Oxford Street",
		Email:     "info@subway.co.uk",
		CreatedAt: time.Now(),
	}
	if err := repo.Update(r); err != nil {
		t.Error(err)
	}
}

func TestDeleteRestaurant(t *testing.T) {
	repo := NewRestaurantRepository(pool)
	if err := repo.Delete(4); err != nil {
		t.Error(err)
	}
}

func TestOpenRestaurant(t *testing.T) {
	var open *bool
	var err error
	repo := NewRestaurantRepository(pool)
	if open, err = repo.OpenClose(5); err != nil {
		t.Error(err)
	}
	if *open != false {
		t.Error("didn't match")
	}
}
