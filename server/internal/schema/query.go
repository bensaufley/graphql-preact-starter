package schema

import (
	"context"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
	"github.com/bensaufley/graphql-preact-starter/internal/log"
	"github.com/bensaufley/graphql-preact-starter/internal/resolvers"
)

func (r *Resolver) GetTodos(ctx context.Context) ([]resolvers.TodoResolver, error) {
	var todos []entities.Todo
	if err := r.DB.WithContext(ctx).Find(&todos); err.Error != nil {
		log.Logger.WithError(err.Error).Error("error finding todos")
		return []resolvers.TodoResolver{}, err.Error
	}
	rslvrs := make([]resolvers.TodoResolver, len(todos))
	for i, t := range todos {
		todo := t
		rslvrs[i] = resolvers.TodoResolver{Todo: todo}
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
