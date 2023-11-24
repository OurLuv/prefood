package postgres

import (
	"testing"
	"time"

	"github.com/OurLuv/prefood/internal/model"
)

func TestCreateOrder(t *testing.T) {
	repo := NewOrderRepository(pool)
	order := model.Order{
		RestaurantId: 3,
		FoodOrder: []model.FoodOrder{
			{Id: 4, Quantity: 1},
			{Id: 2, Quantity: 2},
		},
		Name:     "john doe",
		Phone:    "1234567890",
		Status:   "delivered",
		Channel:  "mobile app",
		Additive: "",
		Discount: 0,
		Ordered:  time.Now(),
		Arrive:   time.Now().Add(time.Minute * 30),
	}
	if err := repo.Create(order); err != nil {
		t.Error(err)
	}

}

func TestGetAllOrders(t *testing.T) {
	repo := NewOrderRepository(pool)

	if orders, err := repo.GetAll(1); err != nil {
		t.Error(err)
		_ = orders
	}
}

func TestGetOrderById(t *testing.T) {
	repo := NewOrderRepository(pool)

	if order, err := repo.GetById(1, 9); err != nil {
		t.Error(err)
		_ = order
	}
}
