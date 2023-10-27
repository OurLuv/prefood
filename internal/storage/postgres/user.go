package postgres

import (
	"context"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage interface {
	Create(f model.User) error
	GetById(id uint) (*model.User, error)
	Login(email string) (*model.User, error)
	UpdateById(id uint, c model.User) error
	DeleteById(id uint) error
}

type UserRepository struct {
	pool *pgxpool.Pool
}

// * Create
func (ur *UserRepository) Create(f model.User) error {
	query := "INSERT INTO users (firstname, lastname, password, email) VALUES ($1, $2, $3, $4)"
	if _, err := ur.pool.Exec(context.Background(), query, f.Firstname, f.Lastname, f.Password, f.Email); err != nil {
		return err
	}
	return nil
}

// * Get by id
func (ur *UserRepository) GetById(id uint) (*model.User, error) {
	query := "SELECT * FROM users WHERE id=$1"
	user := model.User{}
	row := ur.pool.QueryRow(context.Background(), query, id)
	if err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Password, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

// * Login
func (ur *UserRepository) Login(email string) (*model.User, error) {
	query := "SELECT * FROM users where email = $1"

	user := model.User{}
	row := ur.pool.QueryRow(context.Background(), query, email)
	if err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Password, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *UserRepository) UpdateById(id uint, c model.User) error {
	return nil
}
func (ur *UserRepository) DeleteById(id uint) error {
	return nil
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
