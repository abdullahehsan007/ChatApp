package db

import (
	"chatapp/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUser(email string) (bool, error)
	CreateUser(user model.Info) error
	GetId(email string) (string, error)
	GetUserData(email string) (string, string, error)
	GetUserByID(id string) (bool, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}
