package postgres

import (
	"context"
	"fmt"

	"github.com/OurLuv/prefood/internal/common"
	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryStorage interface {
	Create(c model.Category) (uint, error)
	GetById(id uint, restaurantId uint) (*model.Category, error)
	GetAll(id uint) ([]model.Category, error)
	UpdateCategory(c model.Category) error
	DeleteCategoryById(id uint) error
}

type CategoryRepository struct {
	pool *pgxpool.Pool
}

// * Create
func (cr *CategoryRepository) Create(c model.Category) (uint, error) {
	var id uint
	ctx := context.Background()
	row := cr.pool.QueryRow(ctx, "INSERT INTO category (name, restaurant_id) VALUES ($1, $2) RETURNING id", c.Name, c.RestaurantId)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to create a category: %s", err)
	}
	return id, nil
}

// * Get category by id
func (cr *CategoryRepository) GetById(id uint, restaurantId uint) (*model.Category, error) {
	query := "SELECT * FROM category WHERE id = $1 and restaurant_id = $2"
	row := cr.pool.QueryRow(context.Background(), query, id, restaurantId)
	category := &model.Category{}
	err := row.Scan(&category.Id, &category.Name, &category.RestaurantId)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, common.RowNotFound
		}
		return nil, err
	}

	return category, nil
}

// * get all categories
func (cr *CategoryRepository) GetAll(id uint) ([]model.Category, error) {
	query := "SELECT * FROM category WHERE restaurant_id = $1"
	row, err := cr.pool.Query(context.Background(), query, id)
	res := []model.Category{}
	if err != nil {
		return nil, err
	}
	category := model.Category{}
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
func (cr *CategoryRepository) UpdateCategory(c model.Category) error {
	query := "UPDATE category " +
		"SET " +
		"name = $1 " +
		"WHERE id = $2"
	tag, err := cr.pool.Exec(context.Background(), query, c.Name, c.Id)
	if tag.RowsAffected() != 1 {
		return common.RowNotFound
	}
	if err != nil {
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
