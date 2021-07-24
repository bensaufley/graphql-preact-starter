package resolver

import (
	"context"
	"errors"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type TodoInput struct {
	Contents string
	Status   *string
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

func publishTodos(r *Resolver) {
	todos, err := r.GetTodos(context.Background())
	if err != nil {
		log.WithError(err).Warn("could not query todos")
		return
	}
	r.Subscriptions.updateTodos <- &todos
}

func (r *Resolver) AddTodo(ctx context.Context, args *struct{ Input TodoInput }) (*Todo, error) {
	status, err := TodoStatusFromString(args.Input.Status)
	if err != nil {
		return nil, err
	}
	stmt, err := r.DB.PrepareContext(
		ctx,
		`INSERT INTO todos (contents, todo_status_id) VALUES (?, (SELECT id FROM todo_statuses WHERE val = ?)) RETURNING id;`,
	)
	if err != nil {
		return nil, err
	}
	var id int
	if err = stmt.QueryRowContext(ctx, args.Input.Contents, status).Scan(&id); err != nil {
		return nil, err
	}

	go publishTodos(r)

	return &Todo{
		id:       id,
		Contents: args.Input.Contents,
		status:   status,
	}, nil
}

func (r *Resolver) DeleteTodo(ctx context.Context, args *struct{ ID string }) (bool, error) {
	id, err := strconv.Atoi(args.ID)
	if err != nil {
		return false, err
	}
	_, err = r.DB.ExecContext(ctx, `DELETE FROM todos WHERE id = ?;`, id)
	if err != nil {
		return false, err
	}

	go publishTodos(r)

	return true, nil
}
