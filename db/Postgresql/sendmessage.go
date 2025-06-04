package postgresql

import (
	"chatapp/model"
	"time"
)

func (r *userRepo) SendMessage(message model.Message, id string) (string, error) {
	query := `INSERT INTO send (senderid, receiverid, message, time) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, id, message.Receiverid, message.Message, time.Now())
	if err != nil {
		return "", err
	} else {
		return "Message Sent Successfully", nil
	}
}
