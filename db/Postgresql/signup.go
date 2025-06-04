package postgresql

import (
	"chatapp/model"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(user model.Info) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (r *userRepo) GetUser(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM signup WHERE email = $1)`
	err := r.db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func (r *userRepo) CreateUser(user model.Info) error {
	hashed, err := HashedPassword(user)
	if err != nil {
		return err
	}
	query := `INSERT INTO signup(username,email,password) VALUES($1,$2,$3)`
	_, err = r.db.Exec(query, user.Username, user.Email, hashed)
	return err
}
