package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	CategoryStorage
	FoodStorage
	UserStorage
	RestaurantStorage
}

func NewRepository(p *pgxpool.Pool) *Repository {
	return &Repository{
		NewCategoryStorage(p),
		NewFoodRepository(p),
		NewUserRepository(p),
		NewRestaurantRepository(p),
	}
}

func NewDB(storagePath string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), storagePath)
	if err != nil {
		err = fmt.Errorf("unable to create connection pool: %v", err)
		return nil, err
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world from db!'").Scan(&greeting)
	if err != nil {
		err = fmt.Errorf("queryRow failed: %v", err)
		return nil, err
	}
	fmt.Println(greeting)
	return dbpool, nil
}
