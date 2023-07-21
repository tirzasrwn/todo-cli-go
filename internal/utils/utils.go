package utils

import (
	"time"
	"todo-cli-go/internal/constants"
)

func TimeNilChecker(t *time.Time) string {
	if t == nil {
		return "NULL"
	}
	if t.IsZero() {
		return "NULL"
	}
	return t.Format(constants.TimeFormat)
}
