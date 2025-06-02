package db

import "chatapp/model"


func (r *userRepo) GetUserByID(id string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM signup WHERE id = $1)`
	err := r.db.QueryRow(query,id ).Scan(&exists)
	return exists, err
}

func (r *userRepo) SendMessage(message model.Message) (string,err){
	
}