package schema

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
)

type TodoInput struct {
	Contents string
	Status   *string
}

func (r *Resolver) AddTodo(
	ctx context.Context,
	args *struct{ Input TodoInput },
) (*NullableTodoResolver, error) {
	status, err := entities.TodoStatusFromString(args.Input.Status)
	if err != nil {
		return nil, err
	}
	todo := entities.Todo{Contents: args.Input.Contents, Status: status}
	result := r.DB.WithContext(ctx).Create(&todo)
	if result.Error != nil {
		return nil, fmt.Errorf("could not insert Todo: %w", result.Error)
	}

	go r.todoAdded(todo)

	return &NullableTodoResolver{Resolver: r, Todo: todo}, nil
}

func (r *Resolver) AdvanceTodo(
	ctx context.Context,
	args *struct{ ID graphql.ID },
) (*NullableTodoResolver, error) {
	db := r.DB.WithContext(ctx)
	var todo entities.Todo
	if result := db.First(&todo, args.ID); result.Error != nil {
		return nil, fmt.Errorf("could not retrieve Todo: %w", result.Error)
	}
	if err := todo.Status.Advance(); err != nil {
		return nil, err
	}
	if result := db.Save(&todo); result.Error != nil {
		// NB: this is an example app; in practice, you should not return db errors directly
		// to the client because at best it's unhelpful and at worst it's insecure
		return nil, fmt.Errorf("could not update Todo: %w", result.Error)
	}

	go r.todoUpdated(todo)

	return &NullableTodoResolver{Resolver: r, Todo: todo}, nil
}

func (r *Resolver) TransitionTodo(
	ctx context.Context,
	args *struct {
		ID     graphql.ID
		Status string
	},
) (*NullableTodoResolver, error) {
	status, err := entities.TodoStatusFromString(&args.Status)
	if err != nil {
		return nil, fmt.Errorf("could not interpret Status: %w", err)
	}
	var todo entities.Todo
	if result := r.DB.WithContext(ctx).
		Model(&entities.Todo{}).
		First(&todo, args.ID).
		Update("status", status); result.Error != nil {
		return nil, fmt.Errorf("could not retrieve Todo: %w", result.Error)
	}

	todo.Status = status

	go r.todoUpdated(todo)

	return &NullableTodoResolver{Resolver: r, Todo: todo}, nil
}

func (r *Resolver) DeleteTodo(ctx context.Context, args *struct{ ID graphql.ID }) (bool, error) {
	result := r.DB.WithContext(ctx).Delete(&entities.Todo{Base: entities.Base{ULID: string(args.ID)}})
	if result.Error != nil {
		return false, fmt.Errorf("could not delete Todo: %w", result.Error)
	}

	go r.todoDeleted(args.ID)

	return true, nil
}
