package postgres

import (
	"fmt"
	"testing"
	"time"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5"
)

func TestCreate(t *testing.T) {
	fr := NewFoodRepository(pool)
	category := model.Сategory{
		//Id:   25,
		Name: "Fruits",
	}

	food := model.Food{
		Name:        "Apple",
		CategoryId:  category.Id,
		Category:    category,
		Description: "Fresh and juicy apple",
		Price:       10,
		InStock:     true,
		CreatedAt:   time.Now(),
		Image:       "apple.jpg",
	}
	if err := fr.Create(food); err != nil {
		t.Errorf("can't create a model of Food: %v", err)
	}

}

func TestGetById(t *testing.T) {
	fr := NewFoodRepository(pool)
	var err error
	if _, err = fr.GetById(1); err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No rows in table you asked for")
		} else {
			t.Error(err)
		}
	}
}

func TestGetAll(t *testing.T) {
	fr := NewFoodRepository(pool)
	var err error
	var food []model.Food
	if food, err = fr.GetAll(); err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println(food)
		} else {
			t.Error(err)
		}
	}
}
