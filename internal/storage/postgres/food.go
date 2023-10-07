package postgres

import (
	"context"
	"fmt"

	"github.com/OurLuv/prefood/internal/model"
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
	row := fr.pool.QueryRow(ctx, "SELECT * FROM category WHERE id=$1", f.Category.Id)
	if row == nil {
		return fmt.Errorf("category is nil")
	}

	_, err := fr.pool.Exec(ctx, "INSERT INTO food (name, description, category_id, price, in_stock, created_at, image) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		f.Name, f.Description, f.Category.Id, f.Price, f.InStock, f.CreatedAt, f.Image)
	if err != nil {
		return fmt.Errorf("failed to create a category: %s", err)
	}
	return nil
}
