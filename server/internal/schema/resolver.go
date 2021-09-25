package schema

import (
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"

	"github.com/bensaufley/graphql-preact-starter/internal/resolvers"
)

type Resolver struct {
	DB            *gorm.DB
	Subscriptions Subscriptions
}

func NewRoot(db *gorm.DB) *Resolver {
	r := &Resolver{
		DB: db,
		Subscriptions: Subscriptions{
			todoAdded:   make(chan resolvers.TodoResolver),
			todoDeleted: make(chan graphql.ID),

			addedSubscribers:   make(chan *addedSubscriber),
			deletedSubscribers: make(chan *deletedSubscriber),
		},
	}

	go r.broadcastTodoChanges()

	return r
}
