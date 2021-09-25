package entities

import (
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

func (t TodoStatus) ToGQLEnum() string {
	switch t {
	case InProgress:
		return "IN_PROGRESS"
	default:
		return strings.ToUpper(string(t))
	}
}

func (t Todo) StatusGQLEnum() string {
	return t.status.ToGQLEnum()
}

type Todo struct {
	Base
	Contents string
	status   TodoStatus
}

func (t Todo) WithStatus(s TodoStatus) Todo {
	tt := t
	tt.status = s
	return tt
}
