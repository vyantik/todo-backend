package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vyantik/todo-backend"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)

	var itemId int
	row := tx.QueryRow(query, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(query, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	var items []todo.TodoItem
	err := r.db.Select(&items, query, userId, listId)
	return items, err
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND ti.id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	var item todo.TodoItem
	err := r.db.Get(&item, query, userId, itemId)
	return item, err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateItemInput) error {
	query := fmt.Sprintf("UPDATE %s ti SET ", todoItemsTable)
	args := make([]any, 0)
	argId := 1

	if input.Title != nil {
		query += fmt.Sprintf("title = $%d, ", argId)
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		query += fmt.Sprintf("description = $%d, ", argId)
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		query += fmt.Sprintf("done = $%d, ", argId)
		args = append(args, *input.Done)
		argId++
	}

	query = query[:len(query)-2]
	query += fmt.Sprintf(" FROM %s li INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $%d AND ti.id = $%d",
		listsItemsTable, usersListsTable, argId, argId+1)

	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND ti.id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, itemId)
	return err
}
