package postgres

import (
	"context"
	"errors"

	"github.com/OurLuv/prefood/internal/common"
	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderStorage interface {
	Create(order model.Order) (*model.Order, error)
	GetAll(restaurantId uint) ([]model.Order, error)
	GetById(restaurantId uint, order_Id uint) (*model.Order, error)
	Delete(id uint) error
	ChangeStatus(id uint, status string) (string, error)
}

type OrderRepository struct {
	pool *pgxpool.Pool
}

// * Create
func (or *OrderRepository) Create(order model.Order) (*model.Order, error) {
	ctx := context.Background()
	var order_id uint
	tx, err := or.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	// add row to orders
	row := or.pool.QueryRow(ctx, "INSERT INTO orders "+
		"(restaurant_id, name, phone, discount, status, channel, additive, ordered, arrive) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		order.RestaurantId, order.Name, order.Phone, order.Code, order.Status, order.Channel, order.Additive, order.Ordered, order.Arrive)

	if err := row.Scan(&order_id); err != nil {
		return nil, err
	}

	// add rows to order_food
	for f := range order.FoodOrder {
		if _, err := or.pool.Exec(ctx, "INSERT INTO order_food"+
			"(order_id, food_id, quantity) VALUES ($1, $2, $3)", order_id, order.FoodOrder[f].Id, order.FoodOrder[f].Quantity); err != nil {
			return nil, err
		}
	}

	// set total for order
	row = or.pool.QueryRow(ctx, "UPDATE orders "+
		"SET total = ( "+
		"	SELECT SUM(of.quantity * f.price) "+
		"	FROM order_food AS of "+
		"	JOIN food AS f ON of.food_id = f.id "+
		"	WHERE of.order_id = $1 "+
		") "+
		"WHERE orders.id = $1 "+
		"RETURNING *", order_id)
	var ord model.Order
	if err := row.Scan(&ord.Id, &ord.RestaurantId, &ord.Name, &ord.Phone, &ord.Total, &ord.Status, &ord.Channel, &ord.Additive, &ord.Code, &ord.Ordered, &ord.Arrive); err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return &ord, nil
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
		if err := rows.Scan(&order.Id, &order.RestaurantId, &order.Name, &order.Phone, &order.Total, &order.Status, &order.Channel, &order.Additive, &order.Code, &order.Ordered, &order.Arrive); err != nil {
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
	if err := row.Scan(&order.Id, &order.RestaurantId, &order.Name, &order.Phone, &order.Total, &order.Status, &order.Channel, &order.Additive, &order.Code, &order.Ordered, &order.Arrive); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.RowNotFound
		}
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

func (or *OrderRepository) Delete(id uint) error {
	query := "DELETE FROM orders WHERE id = $1"
	if _, err := or.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
func (or *OrderRepository) ChangeStatus(id uint, status string) (string, error) {
	query := "UPDATE orders SET status = $1 WHERE id = $2"
	if _, err := or.pool.Exec(context.Background(), query, status, id); err != nil {
		return "", err
	}
	return status, nil
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		pool: pool,
	}
}
