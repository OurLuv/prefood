package postgres

import (
	"context"
	"fmt"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryStorage interface {
	Create(c model.Сategory) error
	GetById(id uint) (*model.Сategory, error)
	GetAll(id uint) ([]model.Сategory, error)
	UpdateCategory(c model.Сategory) error
	DeleteCategoryById(id uint) error
}

type CategoryRepository struct {
	pool *pgxpool.Pool
}

// * Create
func (cr *CategoryRepository) Create(c model.Сategory) error {
	ctx := context.Background()
	_, err := cr.pool.Exec(ctx, "INSERT INTO category (name, restaurant_id) VALUES ($1, $2)", c.Name, c.RestaurantId)
	if err != nil {
		return fmt.Errorf("failed to create a category: %s", err)
	}
	return nil
}

// * Get category by id
func (cr *CategoryRepository) GetById(id uint) (*model.Сategory, error) {
	query := "SELECT * FROM category WHERE id = $1"
	row := cr.pool.QueryRow(context.Background(), query, id)
	category := &model.Сategory{}
	err := row.Scan(&category.Id, &category.Name, &category.RestaurantId)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// * get all categories
func (cr *CategoryRepository) GetAll(id uint) ([]model.Сategory, error) {
	query := "SELECT * FROM category WHERE restaurant_id = $1"
	row, err := cr.pool.Query(context.Background(), query, id)
	res := []model.Сategory{}
	if err != nil {
		return nil, err
	}
	category := model.Сategory{}
	for row.Next() {
		err := row.Scan(&category.Id, &category.Name, &category.RestaurantId)
		res = append(res, category)
		if err != nil {
			return nil, err
		}

	}
	return res, nil
}

// * update category
func (cr *CategoryRepository) UpdateCategory(c model.Сategory) error {
	query := "UPDATE category " +
		"SET " +
		"name = $1 " +
		"WHERE id = $2"
	if _, err := cr.pool.Exec(context.Background(), query, c.Name, c.Id); err != nil {
		return err
	}
	return nil
}

// * delete category by id
func (cr *CategoryRepository) DeleteCategoryById(id uint) error {
	query := "delete from category where id=$1"
	commandTag, err := cr.pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("no row found to delete")
	}
	return nil
}

func NewCategoryStorage(p *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{pool: p}
}
