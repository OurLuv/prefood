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
	Update(r model.Restaurant) error
	Delete(id uint) error
	OpenClose(id uint) (*bool, error)
}

type RestaurantRepository struct {
	pool *pgxpool.Pool
}

// * Get all
func (rr *RestaurantRepository) GetAll(client_id uint) ([]model.Restaurant, error) {
	query := "SELECT * FROM restaurants WHERE client_id = $1"
	var restaurants []model.Restaurant
	r := model.Restaurant{}
	rows, err := rr.pool.Query(context.Background(), query, client_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&r.Id, &r.Name, &r.ClientId, &r.Phone, &r.Country, &r.State, &r.City, &r.Street, &r.Email, &r.CreatedAt, &r.Open); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, r)
	}
	return restaurants, nil
}

// * Create
func (rr *RestaurantRepository) Create(r model.Restaurant) error {
	query := "INSERT INTO restaurants " +
		"(name, client_id, phone, country, state, city, street, email, created_at, open) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	if _, err := rr.pool.Exec(context.Background(), query,
		r.Name, r.ClientId, r.Phone, r.Country, r.State, r.City, r.Street, r.Email, r.CreatedAt, r.Open); err != nil {
		return err
	}
	return nil
}

// * Get by id
func (rr *RestaurantRepository) GetById(id uint, client_id uint) (*model.Restaurant, error) {
	query := "SELECT * FROM restaurants WHERE id = $1 AND client_id =$2"
	row := rr.pool.QueryRow(context.Background(), query, id, client_id)
	var r model.Restaurant
	if err := row.Scan(&r.Id, &r.Name, &r.ClientId, &r.Phone, &r.Country, &r.State, &r.City, &r.Street, &r.Email, &r.CreatedAt, &r.Open); err != nil {
		return nil, err
	}
	return &r, nil
}

// * Update
func (rr *RestaurantRepository) Update(r model.Restaurant) error {
	query := "UPDATE restaurants " +
		"SET " +
		"name = $1, " +
		"phone = $2, " +
		"country = $3, " +
		"state = $4, " +
		"city = $5, " +
		"street = $6 " +
		"WHERE id = $7"
	if _, err := rr.pool.Exec(context.Background(), query, r.Name, r.Phone, r.Country, r.State, r.City, r.Street, r.Id); err != nil {
		return err
	}
	return nil
}

// * Delete by id
func (rr *RestaurantRepository) Delete(id uint) error {
	query := "DELETE FROM restaurants WHERE id = $1"
	if _, err := rr.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}

// * Open
func (rr *RestaurantRepository) OpenClose(id uint) (*bool, error) {
	var open bool
	query := `UPDATE restaurants SET "open" = NOT "open" WHERE id = $1 RETURNING "open"`
	row := rr.pool.QueryRow(context.Background(), query, id)
	if err := row.Scan(&open); err != nil {
		return nil, err
	}
	return &open, nil
}

func NewRestaurantRepository(pool *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{
		pool: pool,
	}
}
