package postgres

import (
	"context"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RestaurantStorage interface {
	GetAll(client_id uint) ([]model.Restaurant, error)
	Create(r model.Restaurant) error
	GetById(id uint, client_id uint) (*model.Restaurant, error)
}

type RestaurantRepository struct {
	pool *pgxpool.Pool
}

func (rr *RestaurantRepository) GetAll(client_id uint) ([]model.Restaurant, error) {
	query := "SELECT * FROM restaurants WHERE client_id = $1"
	var restaurants []model.Restaurant
	var r model.Restaurant
	rows, err := rr.pool.Query(context.Background(), query, client_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&r.Id, &r.Name, &r.ClientId, &r.Phone, &r.Country, &r.State, &r.City, &r.Street, &r.Email, &r.CreatedAt); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, r)
	}
	return restaurants, nil
}

// * Create
func (rr *RestaurantRepository) Create(r model.Restaurant) error {
	query := "INSERT INTO restaurants " +
		"(name, client_id, phone, country, state, city, street, email, created_at) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	if _, err := rr.pool.Exec(context.Background(), query,
		r.Name, r.ClientId, r.Phone, r.Country, r.State, r.City, r.Street, r.Email, r.CreatedAt); err != nil {
		return err
	}
	return nil
}
func (rr *RestaurantRepository) GetById(id uint, client_id uint) (*model.Restaurant, error) {
	return nil, nil
}

func NewRestaurantRepository(pool *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{
		pool: pool,
	}
}
