package postgres

import (
	"context"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderStorage interface {
	Create(order model.Order) error
	GetAll(restaurantId uint) ([]model.Order, error)
	GetById(restaurantId uint, order_Id uint) (*model.Order, error)
}

type OrderRepository struct {
	pool *pgxpool.Pool
}

// * Create
func (or *OrderRepository) Create(order model.Order) error {
	ctx := context.Background()
	var order_id uint
	tx, err := or.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(context.Background())
	}()
	row := or.pool.QueryRow(ctx, "INSERT INTO orders"+
		"(restaurant_id, name, phone, total, discount, status, channel, additive, ordered, arrive)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		order.RestaurantId, order.Name, order.Phone, order.Total, order.Discount, order.Status, order.Channel, order.Additive, order.Ordered, order.Arrive)

	if err := row.Scan(&order_id); err != nil {
		return err
	}
	for f := range order.FoodOrder {
		if _, err := or.pool.Exec(ctx, "INSERT INTO order_food"+
			"(order_id, food_id, quantity) VALUES ($1, $2, $3)", order_id, order.FoodOrder[f].Id, order.FoodOrder[f].Quantity); err != nil {
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

// * Get all
func (or *OrderRepository) GetAll(restaurantId uint) ([]model.Order, error) {
	query := "SELECT * FROM orders WHERE restaurant_id = $1"
	rows, err := or.pool.Query(context.Background(), query, restaurantId)
	if err != nil {
		return nil, err
	}
	order := model.Order{}
	orderList := []model.Order{}
	for rows.Next() {
		if err := rows.Scan(&order.Id, &order.RestaurantId, &order.Name, &order.Phone, &order.Total, &order.Status, &order.Channel, &order.Additive, &order.Discount, &order.Ordered, &order.Arrive); err != nil {
			return nil, err
		}
		orderList = append(orderList, order)
	}
	//adding food in order
	query = "SELECT f.id, f.name, f.description, f.category_id, f.price, f.in_stock, f.created_at, f.image, of.order_id, of.food_id, of.quantity FROM food f " +
		"LEFT JOIN order_food of ON f.id = of.food_id  " +
		"LEFT JOIN orders o ON of.order_id = o.id  " +
		"WHERE o.restaurant_id = $1 " +
		"ORDER BY of.order_id"
	rows, err = or.pool.Query(context.Background(), query, restaurantId)
	if err != nil {
		return nil, err
	}
	var food model.Food
	var foodOrder model.FoodOrder
	var order_id int
	i := 0
	for rows.Next() {
		if err := rows.Scan(&food.Id, &food.Name, &food.Description, &food.CategoryId, &food.Price, &food.InStock, &food.CreatedAt, &food.Image, &order_id, &foodOrder.Id, &foodOrder.Quantity); err != nil {
			return nil, err
		}
		for i < len(orderList) {
			if order_id == int(orderList[i].Id) {
				orderList[i].Food = append(orderList[i].Food, food)
				orderList[i].FoodOrder = append(orderList[i].FoodOrder, foodOrder)
				break
			} else {
				i++
			}
		}
	}
	return orderList, nil
}

// * Get by id
func (or *OrderRepository) GetById(restaurantId uint, orderId uint) (*model.Order, error) {
	query := "SELECT * FROM orders WHERE restaurant_id = $1 AND id=$2"
	row := or.pool.QueryRow(context.Background(), query, restaurantId, orderId)
	order := model.Order{}
	if err := row.Scan(&order.Id, &order.RestaurantId, &order.Name, &order.Phone, &order.Total, &order.Status, &order.Channel, &order.Additive, &order.Discount, &order.Ordered, &order.Arrive); err != nil {
		return nil, err
	}
	query = "SELECT f.id, f.name, f.description, f.category_id, f.price, f.in_stock, f.created_at, f.image, of.order_id, of.food_id, of.quantity FROM food f " +
		"LEFT JOIN order_food of ON f.id = of.food_id  " +
		"LEFT JOIN orders o ON of.order_id = o.id  " +
		"WHERE o.restaurant_id = $1 AND o.id = $2 " +
		"ORDER BY of.order_id"
	rows, err := or.pool.Query(context.Background(), query, restaurantId, orderId)
	if err != nil {
		return nil, err
	}
	var food model.Food
	var foodOrder model.FoodOrder
	var order_id int
	for rows.Next() {
		if err := rows.Scan(&food.Id, &food.Name, &food.Description, &food.CategoryId, &food.Price, &food.InStock, &food.CreatedAt, &food.Image, &order_id, &foodOrder.Id, &foodOrder.Quantity); err != nil {
			return nil, err
		}
		order.Food = append(order.Food, food)
		order.FoodOrder = append(order.FoodOrder, foodOrder)
	}
	return &order, nil
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		pool: pool,
	}
}
