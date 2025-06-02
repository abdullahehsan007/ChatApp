package db

func (r *userRepo) GetId(email string) (string, error) {
	var id string
	query := `SELECT id FROM signup WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
func (r *userRepo) GetUserData(email string) (string, string, error) {
	var id string
	var dbpass string
	query := `SELECT id,password FROM signup WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&id, &dbpass)
	if err != nil {
		return "", "", err
	}
	return id, dbpass, nil
}
