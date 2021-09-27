package entities

import (
	"database/sql/driver"
	"errors"
	"fmt"
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

func (t TodoStatus) String() string {
	return string(t)
}

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

func TodoStatusFromString(str *string) (TodoStatus, error) {
	if str == nil {
		return Unstarted, nil
	}

	switch *str {
	case "UNSTARTED":
		return Unstarted, nil
	case "IN_PROGRESS":
		return InProgress, nil
	case "ABANDONED":
		return Abandoned, nil
	case "COMPLETE":
		return Complete, nil
	case "DELETED":
		return Deleted, nil
	default:
		return Unstarted, errors.New("unrecognized TodoStatus")
	}
}

func (t *TodoStatus) Advance() error {
	if t == nil {
		return errors.New("todoStatus is nil")
	}
	switch *t {
	case Unstarted:
		*t = InProgress
	case InProgress:
		*t = Complete
	default:
		return fmt.Errorf("cannot advance from status \"%s\"", t.String())
	}
	return nil
}
