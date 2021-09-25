package schema

import (
	"context"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
	"github.com/bensaufley/graphql-preact-starter/internal/resolvers"
)

func (r *Resolver) GetTodos(ctx context.Context) ([]resolvers.TodoResolver, error) {
	var todos []entities.Todo
	if err := r.DB.WithContext(ctx).Find(&todos); err.Error != nil {
		return []resolvers.TodoResolver{}, err.Error
	}
	rslvrs := make([]resolvers.TodoResolver, len(todos))
	for _, t := range todos {
		todo := t
		rslvrs = append(rslvrs, resolvers.TodoResolver{Todo: todo})
	}
	return rslvrs, nil
}

func (r *Resolver) GetTodo(ctx context.Context, args *struct{ ID string }) (*resolvers.TodoResolver, error) {
	var todo entities.Todo
	if err := r.DB.WithContext(ctx).First(&todo, "id = ?", args.ID); err.Error != nil {
		return nil, err.Error
	}
	return &resolvers.TodoResolver{Todo: todo}, nil
}
