package postgres

import (
	"testing"

	"github.com/OurLuv/prefood/internal/model"
)

func TestCreateUser(t *testing.T) {
	repo := NewUserRepository(pool)
	u := model.User{
		Firstname: "Andrew",
		Lastname:  "Orlow",
		Password:  "gghh543s",
		Email:     "aorlowde@gmail.com",
	}
	if err := repo.Create(u); err != nil {
		t.Error(err)
	}
}

func TestGetUser(t *testing.T) {
	repo := NewUserRepository(pool)
	u := model.User{
		Id:        1,
		Firstname: "Andrew",
		Lastname:  "Orlow",
		Password:  "gghh543ss",
		Email:     "aorlowde@gmail.com",
	}
	user, err := repo.Login("aorlowde@gmail.com")
	if err != nil {
		t.Error(err)
	}
	if user.Id != u.Id || user.Firstname != u.Firstname || user.Lastname != u.Lastname || user.Password != u.Password || user.Email != u.Email {
		t.Errorf("Expected: %+v, bot got: %+v", u, user)
	}
}
