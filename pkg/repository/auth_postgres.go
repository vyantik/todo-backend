package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vyantik/todo-backend"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	var id int
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	query := fmt.Sprintf("SELECT id, name, username, password_hash FROM %s WHERE username = $1 AND password_hash = $2", usersTable)

	var user todo.User
	row := r.db.QueryRow(query, username, password)
	if err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Password); err != nil {
		return todo.User{}, err
	}

	return user, nil
}
