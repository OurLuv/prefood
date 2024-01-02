package postgres

import (
	"context"
	"time"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FoodStorage interface {
	Create(f model.Food) (*model.Food, error)
	GetById(restaurantId uint, id uint) (*model.Food, error)
	GetAll(id uint) ([]model.Food, error)
	UpdateById(c model.Food) (*model.Food, error)
	DeleteById(id uint) error
}

type FoodRepository struct {
	pool *pgxpool.Pool
}

// * Create
func (fr *FoodRepository) Create(f model.Food) (*model.Food, error) {
	ctx := context.Background()
	var exist bool
	//* check if category exists
	if err := fr.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM category WHERE id = $1)", f.Category.Id).Scan(&exist); err != nil {
		return nil, err
	}
	if !exist {
		//* create transaction
		tx, err := fr.pool.BeginTx(context.Background(), pgx.TxOptions{})
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = tx.Rollback(context.Background())
		}()
		//* create category
		if _, err := tx.Exec(context.Background(), "INSERT INTO category (id, name, restaurant_id) VALUES ($1, $2, $3)", f.Category.Id, f.Category.Name, f.Category.RestaurantId); err != nil {
			return nil, err
		}
		//* create food
		if _, err := fr.pool.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"); err != nil {
			return nil, err
		}
		row := tx.QueryRow(context.Background(), "INSERT INTO food (name, description, category_id, price, in_stock, created_at, image, restaurant_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *",
			f.Name, f.Description, f.Category.Id, f.Price, f.InStock, time.Now(), "uuid_generate_v4() || '"+f.Image+"'", f.RestaurantId)
		if err := row.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image, &f.RestaurantId); err != nil {
			return nil, err
		}

		//* commit
		if err := tx.Commit(context.Background()); err != nil {
			return nil, err
		}
		return &f, nil
	}
	if _, err := fr.pool.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"); err != nil {
		return nil, err
	}
	row := fr.pool.QueryRow(ctx, "INSERT INTO food (name, description, category_id, price, in_stock, created_at, image, restaurant_id) VALUES ($1, $2, $3, $4, $5, $6, uuid_generate_v4() || '.' || $7, $8) RETURNING *",
		f.Name, f.Description, f.Category.Id, f.Price, f.InStock, time.Now(), f.Image, f.RestaurantId)
	if err := row.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image, &f.RestaurantId); err != nil {
		return nil, err
	}
	return &f, nil
}

// * Get by id
func (fr *FoodRepository) GetById(restaurantId uint, id uint) (*model.Food, error) {
	query := "SELECT * FROM food f JOIN category c on f.category_id = c.id WHERE f.restaurant_id = $1 AND f.id = $2"
	var f model.Food
	var c model.Category
	row := fr.pool.QueryRow(context.Background(), query, restaurantId, id)
	if err := row.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image, &f.RestaurantId, &c.Id, &c.Name, &c.RestaurantId); err != nil {
		return nil, err
	}
	f.Category = c
	return &f, nil
}

// * Get all
func (fr *FoodRepository) GetAll(id uint) ([]model.Food, error) {
	query := "SELECT * FROM food f JOIN category c on f.category_id = c.id WHERE f.restaurant_id = $1"
	var f model.Food
	var c model.Category
	var food []model.Food
	if rows, err := fr.pool.Query(context.Background(), query, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			if err := rows.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image, &f.RestaurantId, &c.Id, &c.Name, &c.RestaurantId); err != nil {
				return nil, err
			}
			f.Category = c
			food = append(food, f)
		}

		return food, nil
	}

}

// * update
func (fr *FoodRepository) UpdateById(f model.Food) (*model.Food, error) {
	query := "UPDATE food " +
		"SET " +
		"name = $1, " +
		"category_id = $2, " +
		"description = $3, " +
		"price = $4, " +
		"in_stock = $5, " +
		"image = $6 " +
		"WHERE id = $7 " +
		"RETURNING *"
	row := fr.pool.QueryRow(context.Background(), query, f.Name, f.CategoryId, f.Description, f.Price, f.InStock, f.Image, f.Id)
	if err := row.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image, &f.RestaurantId); err != nil {
		return nil, err
	}
	return &f, nil
}
func (fr *FoodRepository) DeleteById(id uint) error {
	query := "DELETE FROM food WHERE id = $1"
	if _, err := fr.pool.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}

func NewFoodRepository(p *pgxpool.Pool) *FoodRepository {
	return &FoodRepository{pool: p}
}
