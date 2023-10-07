package postgres

import (
	"log"
	"testing"
	"time"

	"github.com/OurLuv/prefood/internal/model"
)

func TestCreate(t *testing.T){
	pool, err := NewDB("postgres://postgres:admin@localhost:5432/prefood")
	if err != nil{
		log.Fatalf("failed to init storage: %d", err)
	}
	defer pool.Close()

	fr := new(FoodRepository)
	fr.pool = pool
	category := model.Сategory{
		Id: 30,
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
	if err := fr.Create(food); err != nil{
		t.Errorf("can't create a model of Food: %d", err)
	}
	
}