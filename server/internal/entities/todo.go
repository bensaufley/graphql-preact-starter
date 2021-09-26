package entities

type Todo struct {
	Base
	Contents string
	Status   TodoStatus
}
