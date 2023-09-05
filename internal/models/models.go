package models

import "database/sql"

type Models struct {
	Users *UserModel
}

func New(db *sql.DB) (*Models, error) {
	var m Models
	var err error

	m.Users, err = NewUserModel(db)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
