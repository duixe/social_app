package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password password `json:"-"`
	CreatedAt string `json:"created_at"`
}

type password struct {
	text *string
	hash []byte
}

func (pass *password) Set (text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	pass.text = &text
	pass.hash = hash

	return nil
}