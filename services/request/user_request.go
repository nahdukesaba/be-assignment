package request

import "errors"

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *UserRequest) Validate() error {
	if s.Username == "" {
		return errors.New("username cannot be empty")
	}

	if s.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}
