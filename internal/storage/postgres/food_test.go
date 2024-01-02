package postgres

import (
	"testing"
	"time"

	"github.com/OurLuv/prefood/internal/model"
)

func TestCreateFood(t *testing.T) {
	fr := NewFoodRepository(pool)
	category := model.Category{
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
	if _, err := fr.Create(food); err != nil {
		t.Errorf("can't create a model of Food: %v", err)
	}

}

func TestUpdateFood(t *testing.T) {
	fr := NewFoodRepository(pool)
	category := model.Category{
		//Id:   25,
		Name: "Fruits",
	}

	food := model.Food{
		Id:          1,
		Name:        "Apple Pie",
		CategoryId:  category.Id,
		Category:    category,
		Description: "Fresh and juicy apple pie",
		Price:       599,
		InStock:     true,
		CreatedAt:   time.Now(),
		Image:       "apple.jpg",
	}
	if _, err := fr.UpdateById(food); err != nil {
		t.Errorf("can't create a model of Food: %v", err)
	}

}

func TestGetById(t *testing.T) {
	// fr := NewFoodRepository(pool)
	// var err error
	// if _, err = fr.GetById(1); err != nil {
	// 	if err == pgx.ErrNoRows {
	// 		fmt.Println("No rows in table you asked for")
	// 	} else {
	// 		t.Error(err)
	// 	}
	// }
}

func TestGetAll(t *testing.T) {
	fr := NewFoodRepository(pool)
	var err error
	var food []model.Food
	if food, err = fr.GetAll(5); err != nil {
		t.Error(err)
	}
	_ = food
}
