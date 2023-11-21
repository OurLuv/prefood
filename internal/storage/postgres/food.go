package postgres

import (
	"context"
	"time"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FoodStorage interface {
	Create(f model.Food) error
	GetById(restaurantId uint, id uint) (*model.Food, error)
	GetAll(id uint) ([]model.Food, error)
	UpdateById(c model.Food) error
	DeleteById(id uint) error
}

type FoodRepository struct {
	pool *pgxpool.Pool
}

// * Create
func (fr *FoodRepository) Create(f model.Food) error {
	ctx := context.Background()
	var exist bool
	//* check if category exists
	if err := fr.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM category WHERE id = $1)", f.Category.Id).Scan(&exist); err != nil {
		return err
	}
	if !exist {
		//* create transaction
		tx, err := fr.pool.BeginTx(context.Background(), pgx.TxOptions{})
		if err != nil {
			return err
		}
		defer func() {
			_ = tx.Rollback(context.Background())
		}()
		//* create category
		if _, err := tx.Exec(context.Background(), "INSERT INTO category (id, name) VALUES ($1, $2)", f.Category.Id, f.Category.Name); err != nil {
			return err
		}
		//* create food
		if _, err := tx.Exec(context.Background(), "INSERT INTO food (name, description, category_id, price, in_stock, created_at, image) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			f.Name, f.Description, f.Category.Id, f.Price, f.InStock, time.Now(), f.Image); err != nil {
			return err
		}
		//* commit
		if err := tx.Commit(context.Background()); err != nil {
			return err
		}
		return nil
	}

	_, err := fr.pool.Exec(ctx, "INSERT INTO food (name, description, category_id, price, in_stock, created_at, image) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		f.Name, f.Description, f.Category.Id, f.Price, f.InStock, time.Now(), f.Image)
	if err != nil {
		return err
	}
	return nil
}

// * Get by id
func (fr *FoodRepository) GetById(restaurantId uint, id uint) (*model.Food, error) {
	query := "SELECT * FROM food f JOIN category c on f.category_id = c.id WHERE f.restaurant_id = $1 AND f.id = $2"
	var f model.Food
	var c model.Сategory
	row := fr.pool.QueryRow(context.Background(), query, restaurantId, id)
	if err := row.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image, &f.RestaurantId, &c.Id, &c.Name); err != nil {
		return nil, err
	}
	f.Category = c
	return &f, nil
}

// * Get all
func (fr *FoodRepository) GetAll(id uint) ([]model.Food, error) {
	query := "SELECT * FROM food f JOIN category c on f.category_id = c.id WHERE f.restaurant_id = $1"
	var f model.Food
	var c model.Сategory
	var food []model.Food
	if rows, err := fr.pool.Query(context.Background(), query, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			if err := rows.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image, &f.RestaurantId, &c.Id, &c.Name); err != nil {
				return nil, err
			}
			f.Category = c
			food = append(food, f)
		}

		return food, nil
	}

}
func (fr *FoodRepository) UpdateById(f model.Food) error {
	query := "UPDATE food " +
		"SET " +
		"name = $1, " +
		"category_id = $2, " +
		"description = $3, " +
		"price = $4, " +
		"in_stock = $5, " +
		"image = $6 " +
		"WHERE id = $7"
	if _, err := fr.pool.Exec(context.Background(), query, f.Name, f.CategoryId, f.Description, f.Price, f.InStock, f.Image, f.Id); err != nil {
		return err
	}
	return nil
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
