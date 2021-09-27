package schema

import (
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type Resolver struct {
	DB            *gorm.DB
	Subscriptions Subscriptions
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
