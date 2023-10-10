package postgres

import (
	"context"
	"fmt"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FoodStorage interface {
	Create(f model.Food) error
	GetById(id uint) (*model.Food, error)
	GetAll() ([]model.Food, error)
	UpdateById(id uint, c model.Food) error
	DeleteById(id uint) error
}

type FoodRepository struct {
	pool *pgxpool.Pool
}

func (fr *FoodRepository) Create(f model.Food) error {
	ctx := context.Background()
	var exist bool
	//* check if category exists
	row := fr.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM food WHERE id = $1)", f.Category.Id).Scan(&exist)
	if row == nil {
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
			f.Name, f.Description, f.Category.Id, f.Price, f.InStock, f.CreatedAt, f.Image); err != nil {
			return err
		}
		//* commit
		if err := tx.Commit(context.Background()); err != nil {
			return err
		}
	}

	_, err := fr.pool.Exec(ctx, "INSERT INTO food (name, description, category_id, price, in_stock, created_at, image) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		f.Name, f.Description, f.Category.Id, f.Price, f.InStock, f.CreatedAt, f.Image)
	if err != nil {
		return fmt.Errorf("failed to create a category: %s", err)
	}
	return nil
}
