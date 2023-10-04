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
	GetAll() ([]model.Сategory, error)
	UpdateCategoryById(id uint, c model.Сategory) error
	DeleteCategoryById(id uint) error
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
func (cr *CategoryRepository) GetById(id uint) (*model.Сategory, error){
	query := "SELECT id, name FROM category WHERE id = $1"
	row := cr.pool.QueryRow(context.Background(), query, id)
	category := &model.Сategory{}
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		return nil, err
	}
	
	return category, nil
}
//* get all categories
func (cr *CategoryRepository) GetAll() ([]model.Сategory, error){
	query := "SELECT * FROM category"
	row, err := cr.pool.Query(context.Background(), query)
	res := []model.Сategory{}
	if err != nil{
		return nil, err
	}
	category := model.Сategory{}
	for row.Next() {
		err := row.Scan(&category.Id, &category.Name)
		res = append(res, category)
		if err != nil {
			return nil, err
		}
		
	}
	return res, nil
}

//* update category
func (cr *CategoryRepository) UpdateCategoryById(id uint, c model.Сategory) error{
	return nil
}

//* delete category by id
func (cr *CategoryRepository) DeleteCategoryById(id uint) error{
	query := "delete from category where id=$1"
	commandTag, err := cr.pool.Exec(context.Background(), query, id)
	if err != nil{
		return err
	}	
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("no row found to delete")
	}
	return nil
}



func NewCategoryStorage (p *pgxpool.Pool) *CategoryRepository{
	return &CategoryRepository{pool: p}
}