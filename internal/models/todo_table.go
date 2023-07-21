package models

import (
	"context"
	"database/sql"
	"time"
	"todo-cli-go/internal/utils"
)

type DBModels struct {
	DB *sql.DB
}

func (m *DBModels) MakeTodoTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
create table if not exists todo(
  id integer primary key autoincrement not null,
  content text not null,
  isDone boolean not null,
  createdAt datetime not null default current_time,
  updatedAt datetime not null default current_time,
  doneAt datetime
);`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModels) GetAllTodos() ([]*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := "select id, content, isDone, createdAt, updatedAt, doneAt from todo"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []*Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Content, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt, &todo.DoneAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (m *DBModels) InsertTodo(content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `insert into todo (content, isDone, createdAt, updatedAt) values ($1, $2, $3, $4)`
	_, err := m.DB.ExecContext(ctx, query, content, false, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModels) IsIdExist(id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select case when count(id) > 0 then 1 else 0 end from todo where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (m *DBModels) GetTodoById(id int64) (*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, content, isDone, createdAt, updatedAt, doneAt from todo where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)
	var t Todo
	err := row.Scan(&t.Id, &t.Content, &t.IsDone, &t.CreatedAt, &t.UpdatedAt, &t.DoneAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (m *DBModels) DeleteTodoById(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `delete from todo where id = $1`
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModels) UpdateContentById(id int64, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `update todo set content = $1, updatedAt = $2 where id = $3`
	_, err := m.DB.ExecContext(ctx, query, content, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModels) ToggleIsDoneById(id int64) error {
	todo, err := m.GetTodoById(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if !todo.IsDone {
		now := time.Now()
		todo.DoneAt = &now
	} else {
		todo.DoneAt = nil
	}
	query := `update todo set isDone = $1, updatedAt = $2, doneAt = $3 where id = $4`
	_, err = m.DB.ExecContext(ctx, query, !todo.IsDone, time.Now(), utils.TimeNilChecker(todo.DoneAt), id)
	if err != nil {
		return err
	}
	return nil
}
