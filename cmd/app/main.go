package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"todo-cli-go/internal/constants"
	"todo-cli-go/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	logger  *log.Logger
	models  models.Models
	command constants.Command
}

func main() {
	var help bool

	flag.BoolVar(&help, "h", false, "Show help message")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.Parse()

	if help {
		ShowHelpMessage()
		return
	}

	args := os.Args[1:]

	if len(args) < 1 {
		ShowHelpMessage()
		return
	}

	logger := log.New(os.Stdout, "--> ", log.Ldate|log.Ltime)

	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		logger.Panic(err)
		return
	}
	defer db.Close()

	app := &application{
		logger:  logger,
		models:  models.NewModels(db),
		command: constants.Command(args[0]),
	}

	err = app.models.DB.MakeTodoTable()
	if err != nil {
		logger.Panic(err)
		return
	}

	switch app.command {
	case constants.ReadAll:
		app.logger.Println("readall")
		err = app.ReadAll()
		if err != nil {
			logger.Panic(err)
			return
		}
	case constants.Create:
		app.logger.Println("create")
		if len(args) < 2 {
			logger.Panic("need more argument for content")
			return
		}
		err = app.CreateTodo(args[1])
		if err != nil {
			logger.Panic(err)
			return
		}
	case constants.Read:
		app.logger.Println("read")
		if len(args) < 2 {
			logger.Panic("need more argument for id")
			return
		}
		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			logger.Panic(err)
			return
		}
		err = app.ReadTodoById(id)
		if err != nil {
			logger.Panic(err)
			return
		}
	case constants.Delete:
		app.logger.Println("delete")
		if len(args) < 2 {
			logger.Panic("need more argument for id")
			return
		}
		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			logger.Panic(err)
			return
		}
		err = app.DeleteTodoById(id)
		if err != nil {
			logger.Panic(err)
			return
		}
	case constants.Update:
		app.logger.Println("update")
		if len(args) < 3 {
			logger.Panic("need more argument for id and content")
			return
		}
		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			logger.Panic(err)
			return
		}
		err = app.UpdateContentById(id, args[2])
		if err != nil {
			logger.Panic(err)
			return
		}
	case constants.Toggle:
		app.logger.Println("toggle")
		if len(args) < 2 {
			logger.Panic("need more argument for id")
			return
		}
		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			logger.Panic(err)
			return
		}
		err = app.ToggleIsDoneById(id)
		if err != nil {
			logger.Panic(err)
			return
		}
	case constants.Done:
		app.logger.Println("done")
		err := app.FilterTask(true)
		if err != nil {
			logger.Panic(err)
			return
		}
	case constants.Undone:
		app.logger.Println("undone")
		err := app.FilterTask(false)
		if err != nil {
			logger.Panic(err)
			return
		}
	default:
		app.logger.Println("command not avaliable")
		ShowHelpMessage()
		return
	}

	logger.Println("command", app.command, "success")
}

func ShowHelpMessage() {
	help := `
  usage: todo [options] command

  commands:
    readall   Show all tasks
    create    Create new task
    read      Show task by id
    update    Update content task by id
    delete    Delete task by id
    toggle    Toggle done task by id
    done      Show done task
    undone    Show undone task
    help      Show this help message

  options:
    -h, --help Show this help message
`
	fmt.Println(help)
}
