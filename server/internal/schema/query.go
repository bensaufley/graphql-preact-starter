package schema

import (
	"context"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
	"github.com/bensaufley/graphql-preact-starter/internal/log"
)

func (r *Resolver) GetTodos(ctx context.Context) ([]TodoResolver, error) {
	var todos []entities.Todo
	if err := r.DB.WithContext(ctx).Find(&todos); err.Error != nil {
		log.Logger.WithError(err.Error).Error("error finding todos")
		return []TodoResolver{}, err.Error
	}
	rslvrs := make([]TodoResolver, len(todos))
	for i, t := range todos {
		todo := t
		rslvrs[i] = TodoResolver{Todo: todo}
	}
	return rslvrs, nil
}

func (r *Resolver) GetTodo(ctx context.Context, args *struct{ ID string }) (*TodoResolver, error) {
	var todo entities.Todo
	if err := r.DB.WithContext(ctx).First(&todo, "id = ?", args.ID); err.Error != nil {
		return nil, err.Error
	}
	return &TodoResolver{Resolver: r, Todo: todo}, nil
}
