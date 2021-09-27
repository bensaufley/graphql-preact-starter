package schema

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"

	"github.com/bensaufley/graphql-preact-starter/internal/ulid"
)

type Subscriptions struct {
	todoAdded   chan TodoResolver
	todoUpdated chan TodoResolver
	todoDeleted chan graphql.ID

	addedSubscribers   chan *todoSubscriber
	updatedSubscribers chan *todoSubscriber
	deletedSubscribers chan *deletedSubscriber
}

type todoSubscriber struct {
	stop   <-chan struct{}
	events chan<- TodoResolver
}

type deletedSubscriber struct {
	stop   <-chan struct{}
	events chan<- graphql.ID
}

func (r *Resolver) TodoAdded(ctx context.Context) (<-chan TodoResolver, error) {
	ch := make(chan TodoResolver)
	r.Subscriptions.addedSubscribers <- &todoSubscriber{events: ch, stop: ctx.Done()}
	return ch, nil
}

func (r *Resolver) TodoUpdated(ctx context.Context) (<-chan TodoResolver, error) {
	ch := make(chan TodoResolver)
	r.Subscriptions.updatedSubscribers <- &todoSubscriber{events: ch, stop: ctx.Done()}
	return ch, nil
}

func (r *Resolver) TodoDeleted(ctx context.Context) (<-chan graphql.ID, error) {
	ch := make(chan graphql.ID)
	r.Subscriptions.deletedSubscribers <- &deletedSubscriber{events: ch, stop: ctx.Done()}
	return ch, nil
}

func (r *Resolver) broadcastTodoChanges() {
	addedSubscribers := map[string]*todoSubscriber{}
	updatedSubscribers := map[string]*todoSubscriber{}
	deletedSubscribers := map[string]*deletedSubscriber{}
	unsubscribe := make(chan string)
	ug := ulid.NewGenerator()

	for {
		select {
		case id := <-unsubscribe:
			delete(addedSubscribers, id)
		case s := <-r.Subscriptions.addedSubscribers:
			addedSubscribers[ug.String()] = s
		case s := <-r.Subscriptions.updatedSubscribers:
			updatedSubscribers[ug.String()] = s
		case s := <-r.Subscriptions.deletedSubscribers:
			deletedSubscribers[ug.String()] = s
		case e := <-r.Subscriptions.todoAdded:
			for id, s := range addedSubscribers {
				go func(id string, s *todoSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		case e := <-r.Subscriptions.todoUpdated:
			for id, s := range updatedSubscribers {
				go func(id string, s *todoSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						case s.events <- e:
						case <-time.After(time.Second):
					}
				}(id, s)
			}
		case e := <-r.Subscriptions.todoDeleted:
			for id, s := range deletedSubscribers {
				go func(id string, s *deletedSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}
