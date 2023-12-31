package main

import (
	"errors"
	"fmt"
	"todo-cli-go/internal/models"
)

func (app *application) ReadAll() error {
	todos, err := app.models.DB.GetAllTodos()
	if err != nil {
		app.logger.Panic(err)
		return err
	}
	var undone, done, total int64
	for _, todo := range todos {
		total++
		if todo.IsDone {
			done++
		} else {
			undone++
		}
	}
	ts := models.Todos(todos)
	ts.Print()
	fmt.Printf("Undone task is %d, done task is %d, and total task is %d\n", undone, done, total)
	return nil
}

func (app *application) CreateTodo(content string) error {
	err := app.models.DB.InsertTodo(content)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) ReadTodoById(id int64) error {
	exist, err := app.models.DB.IsIdExist(id)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("id not found")
	}
	todo, err := app.models.DB.GetTodoById(id)
	if err != nil {
		return err
	}
	todo.Print()
	return nil
}

func (app *application) FilterTask(isDone bool) error {
	todos, err := app.models.DB.GetTodoByIsDone(isDone)
	if err != nil {
		return err
	}
	ts := models.Todos(todos)
	ts.Print()
	return nil
}

func (app *application) DeleteTodoById(id int64) error {
	exist, err := app.models.DB.IsIdExist(id)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("id not found")
	}
	err = app.models.DB.DeleteTodoById(id)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) UpdateContentById(id int64, content string) error {
	exist, err := app.models.DB.IsIdExist(id)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("id not found")
	}
	err = app.models.DB.UpdateContentById(id, content)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) ToggleIsDoneById(id int64) error {
	exist, err := app.models.DB.IsIdExist(id)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("id not found")
	}
	err = app.models.DB.ToggleIsDoneById(id)
	if err != nil {
		return err
	}
	return nil
}
