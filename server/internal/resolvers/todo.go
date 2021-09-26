// Package resolvers provides wrapper structs to entities to enable nested GraphQL
// resolution.
package resolvers

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
)

type TodoResolver struct {
	entities.Todo
}

type NullableTodoResolver TodoResolver

func (t *NullableTodoResolver) Contents() string {
	if t == nil {
		return ""
	}
	return TodoResolver(*t).Contents()
}

func (t TodoResolver) Contents() string {
	return t.Todo.Contents
}

func (t *NullableTodoResolver) ID() graphql.ID {
	if t == nil {
		return ""
	}
	return TodoResolver(*t).ID()
}

func (t TodoResolver) ID() graphql.ID {
	return graphql.ID(t.Todo.ULID)
}

func (t *NullableTodoResolver) Status() string {
	if t == nil {
		return "UNSTARTED"
	}
	return TodoResolver(*t).Status()
}

func (t TodoResolver) Status() string {
	return t.Todo.Status.ToGQLEnum()
}
