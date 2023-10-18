package postgres

import (
	"context"

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
			f.Name, f.Description, f.Category.Id, f.Price, f.InStock, f.CreatedAt, f.Image); err != nil {
			return err
		}
		//* commit
		if err := tx.Commit(context.Background()); err != nil {
			return err
		}
		return nil
	}

	_, err := fr.pool.Exec(ctx, "INSERT INTO food (name, description, category_id, price, in_stock, created_at, image) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		f.Name, f.Description, f.Category.Id, f.Price, f.InStock, f.CreatedAt, f.Image)
	if err != nil {
		return err
	}
	return nil
}

// * Get by id
func (fr *FoodRepository) GetById(id uint) (*model.Food, error) {
	query := "SELECT * FROM food WHERE id=$1"
	var food model.Food
	row := fr.pool.QueryRow(context.Background(), query, id)
	if err := row.Scan(&food.Id, &food.Name, &food.Description, &food.CategoryId, &food.Price, &food.InStock, &food.CreatedAt, &food.Image); err != nil {
		return nil, err
	}
	return &food, nil
}

// * Get all
func (fr *FoodRepository) GetAll() ([]model.Food, error) {
	query := "SELECT * FROM food"
	var f model.Food
	var food []model.Food
	if rows, err := fr.pool.Query(context.Background(), query); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			if err := rows.Scan(&f.Id, &f.Name, &f.Description, &f.CategoryId, &f.Price, &f.InStock, &f.CreatedAt, &f.Image); err != nil {
				return nil, err
			}
			food = append(food, f)
		}
		return food, nil
	}

}
func (fr *FoodRepository) UpdateById(id uint, c model.Food) error {
	return nil
}
func (fr *FoodRepository) DeleteById(id uint) error {
	return nil
}

func NewFoodRepository(p *pgxpool.Pool) *FoodRepository {
	return &FoodRepository{pool: p}
}
