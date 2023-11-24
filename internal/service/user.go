package service

import (
	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/storage/postgres"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo postgres.UserStorage
}

func (us *UserServiceImpl) Create(u model.User) error {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(passwordByte)
	return us.repo.Create(u)
}
func (us *UserServiceImpl) GetById(id uint) (*model.User, error) {
	return nil, nil
}
func (us *UserServiceImpl) Login(email string, password string) (*model.User, error) {
	var user *model.User
	var err error
	if user, err = us.repo.Login(email); err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
func (us *UserServiceImpl) UpdateById(id uint, c model.User) error {
	return nil
}
func (us *UserServiceImpl) DeleteById(id uint) error {
	return nil
}
func (us *UserServiceImpl) CheckForEmail(email string) error {
	return us.repo.CheckForEmail(email)
}

func NewUserServiceImpl(repo postgres.UserStorage) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}
