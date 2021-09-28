package schema

import (
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
)

type Resolver struct {
	DB            *gorm.DB
	Subscriptions Subscriptions
}

func (r *Resolver) todoAdded(t entities.Todo) {
	r.Subscriptions.todoAdded <- TodoResolver{Resolver: r, Todo: t}
}

func (r *Resolver) todoUpdated(t entities.Todo) {
	r.Subscriptions.todoUpdated <- TodoResolver{Resolver: r, Todo: t}
}

func (r *Resolver) todoDeleted(id graphql.ID) {
	r.Subscriptions.todoDeleted <- id
}

func NewRoot(db *gorm.DB) *Resolver {
	r := &Resolver{
		DB: db,
		Subscriptions: Subscriptions{
			todoAdded:   make(chan TodoResolver),
			todoUpdated: make(chan TodoResolver),
			todoDeleted: make(chan graphql.ID),

			addedSubscribers:   make(chan *todoSubscriber),
			updatedSubscribers: make(chan *todoSubscriber),
			deletedSubscribers: make(chan *deletedSubscriber),
		},
	}

	go r.broadcastTodoChanges()

	return r
}
