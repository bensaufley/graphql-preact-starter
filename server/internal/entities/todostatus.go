package entities

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type TodoStatus string

const (
	Unstarted  TodoStatus = "Unstarted"
	InProgress TodoStatus = "In Progress"
	Abandoned  TodoStatus = "Abandoned"
	Complete   TodoStatus = "Complete"
	Deleted    TodoStatus = "Deleted"
)

func (t *TodoStatus) Scan(value interface{}) error {
	if str, ok := value.(string); ok {
		*t = TodoStatus(str)
		return nil
	}
	return errors.New("could not coerce TodoStatus from database to string")
}

func (t *TodoStatus) Value() (driver.Value, error) {
	return string(*t), nil
}

func (t TodoStatus) ToGQLEnum() string {
	switch t {
	case InProgress:
		return "IN_PROGRESS"
	default:
		return strings.ToUpper(string(t))
	}
}

