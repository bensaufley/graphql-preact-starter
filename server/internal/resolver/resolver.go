package resolver

import "database/sql"

type Resolver struct {
	DB            *sql.DB
	Subscriptions Subscriptions
}

func NewRoot(db *sql.DB) *Resolver {
	r := &Resolver{
		DB: db,
		Subscriptions: Subscriptions{
			updateTodos:     make(chan *[]Todo),
			todoSubscribers: make(chan *todosSubscriber),
		},
	}

	go r.broadcastTodoUpdate()

	return r
}
