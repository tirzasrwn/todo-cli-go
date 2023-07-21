package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"todo-cli-go/internal/constants"
	"todo-cli-go/internal/utils"
)

type Models struct {
	DB DBModels
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModels{DB: db},
	}
}

type Todo struct {
	Id        int64
	Content   string
	IsDone    bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DoneAt    *time.Time
}

type Todos []*Todo

var header string = fmt.Sprintf("|%-3s|%-30s|%-5s|%-25s|%-25s|%-25s|\n",
	"ID", "CONTENT", "DONE", "CREATED_AT", "UPDATED_AT", "DONE_AT")

func (ts *Todos) Print() {
	dash := strings.Repeat("-", len(header)-1)
	fmt.Println(dash)
	fmt.Print(header)
	fmt.Println(dash)
	for _, t := range *ts {
		fmt.Printf("|%-3d|%-30s|%-5t|%-25s|%-25s|%-25s|\n",
			t.Id, t.Content, t.IsDone, t.CreatedAt.Format(constants.TimeFormat),
			t.UpdatedAt.Format(constants.TimeFormat), utils.TimeNilChecker(t.DoneAt))
	}
	fmt.Println(dash)
}

func (t *Todo) Print() {
	dash := strings.Repeat("-", len(header)-1)
	fmt.Println(dash)
	fmt.Print(header)
	fmt.Println(dash)
	fmt.Printf("|%-3d|%-30s|%-5t|%-25s|%-25s|%-25s|\n",
		t.Id, t.Content, t.IsDone, t.CreatedAt.Format(constants.TimeFormat),
		t.UpdatedAt.Format(constants.TimeFormat), utils.TimeNilChecker(t.DoneAt))
	fmt.Println(dash)
}
