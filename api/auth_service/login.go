package authservice

import (
	"chatapp/model"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidGmail(email string) bool {
	valid := regexp.MustCompile(`^[^@]+@gmail\.com$`)
	return valid.MatchString(email)
}
func (r *authService) Login(email, password string) (string, error) {
	return r.Authenticator(model.User{
		Email:    email,
		Password: password,
	})
}
func (r *authService) Authenticator(credential model.User) (string, error) {
	exist, err := r.repo.GetUser(credential.Email)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", fmt.Errorf("user not found")
	}
	id, pass, err := r.repo.GetUserData(credential.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(credential.Password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}
	return id, err
}
