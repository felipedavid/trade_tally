package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	"trade_tally/internal/validator"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Version   int64     `json:"-"`
}

func ValidateUser(v *validator.Validator, u *User) {
	v.Check(u.Email != "", "email", "Must be provided")
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 10)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

type UserModel struct {
	DB      *sql.DB
	GetStmt *sql.Stmt
}

func NewUserModel(db *sql.DB) (*UserModel, error) {
	getStmt, err := db.Prepare(``)
	if err != nil {
		return nil, err
	}

	return &UserModel{DB: db, GetStmt: getStmt}, nil
}

func (m *UserModel) Get(userID int64) (*User, error) {
	return nil, nil
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (m *UserModel) GetAll() ([]*User, error) {
	return nil, nil
}

func (m *UserModel) Insert(user *User) error {
	return nil
}

func (m *UserModel) Update(user *User) error {
	return nil
}

func (m *UserModel) Delete(user *User) error {
	return nil
}
