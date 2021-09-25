package entities

type Todo struct {
	Base
	Contents string
	status   TodoStatus
}

func (t Todo) StatusGQLEnum() string {
	return t.status.ToGQLEnum()
}

func (t Todo) WithStatus(s TodoStatus) Todo {
	tt := t
	tt.status = s
	return tt
}
