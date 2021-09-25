package schema

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
	"github.com/bensaufley/graphql-preact-starter/internal/resolvers"
)

type TodoInput struct {
	Contents string
	Status   *string
}

func TodoStatusFromString(str *string) (entities.TodoStatus, error) {
	if str == nil {
		return entities.Unstarted, nil
	}

	switch *str {
	case "UNSTARTED":
		return entities.Unstarted, nil
	case "IN_PROGRESS":
		return entities.InProgress, nil
	case "ABANDONED":
		return entities.Abandoned, nil
	case "COMPLETE":
		return entities.Complete, nil
	case "DELETED":
		return entities.Deleted, nil
	default:
		return entities.Unstarted, errors.New("unrecognized TodoStatus")
	}
}

func (r *Resolver) AddTodo(
	ctx context.Context,
	args *struct{ Input TodoInput },
) (*resolvers.NullableTodoResolver, error) {
	status, err := TodoStatusFromString(args.Input.Status)
	if err != nil {
		return nil, err
	}
	todo := entities.Todo{Contents: args.Input.Contents}.WithStatus(status)
	result := r.DB.WithContext(ctx).Create(&todo)
	if result.Error != nil {
		return nil, fmt.Errorf("could not insert Todo: %w", result.Error)
	}

	go func(r *Resolver, t resolvers.TodoResolver) {
		r.Subscriptions.todoAdded <- t
	}(r, resolvers.TodoResolver{Todo: todo})

	return &resolvers.NullableTodoResolver{Todo: todo}, nil
}

func (r *Resolver) DeleteTodo(ctx context.Context, args *struct{ ID graphql.ID }) (bool, error) {
	result := r.DB.WithContext(ctx).Delete(&entities.Todo{Base: entities.Base{ULID: string(args.ID)}})
	if result.Error != nil {
		return false, fmt.Errorf("could not delete Todo: %w", result.Error)
	}

	go func(r *Resolver, id graphql.ID) {
		r.Subscriptions.todoDeleted <- id
	}(r, args.ID)

	return true, nil
}
