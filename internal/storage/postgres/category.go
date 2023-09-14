package postgres

import (
	"context"
	"fmt"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryStorage interface {
	Create(c model.Сategory) error
	GetById(c model.Сategory) (*model.Сategory, error)
	GetAll(c model.Сategory) ([]model.Сategory, error)
	Update(c model.Сategory) error
	DeleteById(c model.Сategory) error
}

type CategoryRepository struct{
	pool *pgxpool.Pool
}

//* Create
func (cr *CategoryRepository) Create(c model.Сategory) error{
	ctx := context.Background()
    _, err := cr.pool.Exec(ctx, "INSERT INTO category (name) VALUES ($1)", c.Name)
    if err != nil {
        return fmt.Errorf("failed to create a category: %s", err)
    }
	return nil
}

//* Get category by id
func (cr *CategoryRepository) GetById(c model.Сategory) (*model.Сategory, error){
	return nil, nil
}

//* get all categories
func (cr *CategoryRepository) GetAll(c model.Сategory) ([]model.Сategory, error){
	return nil, nil
}

//* update category
func (cr *CategoryRepository) Update(c model.Сategory) error{
	return nil
}

//* delete category by id
func (cr *CategoryRepository) DeleteById(c model.Сategory) error{
	return nil
}



func NewCategoryStorage (p *pgxpool.Pool) *CategoryRepository{
	return &CategoryRepository{pool: p}
}