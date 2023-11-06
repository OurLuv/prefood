package postgres

import (
	"testing"
	"time"

	"github.com/OurLuv/prefood/internal/model"
)

func TestCreateRestaurant(t *testing.T) {
	repo := NewRestaurantRepository(pool)
	r := model.Restaurant{
		Name:      "Subway",
		ClientId:  3,
		Phone:     "999-888-7777",
		Country:   "United Kingdom",
		State:     "England",
		City:      "London",
		Street:    "654 Oxford Street",
		Email:     "info@subway.co.uk",
		CreatedAt: time.Now(),
	}
	if err := repo.Create(r); err != nil {
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
