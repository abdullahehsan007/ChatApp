package messageservice

import "chatapp/model"

func (s *messageService) GetMessage(get model.Get, Token string) ([]model.Message, error) {
	id, err := DecodeToken(Token)
	if err != nil {
		return nil, err
	}

	messages, err := s.mongorepo.GetMessage(get, id)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
