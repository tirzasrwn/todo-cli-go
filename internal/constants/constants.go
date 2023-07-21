package constants

import "time"

type Command string

const (
	Create  Command = "create"
	ReadAll Command = "readall"
	Read    Command = "read"
	Update  Command = "update"
	Delete  Command = "delete"
	Toggle  Command = "toggle"
)

const TimeFormat string = time.RFC3339
