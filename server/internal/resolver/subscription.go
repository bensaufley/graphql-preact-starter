package resolver

import (
	"context"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
)

type Subscriptions struct {
	updateTodos     chan *[]Todo
	todoSubscribers chan *todosSubscriber
}

type todosSubscriber struct {
	stop   <-chan struct{}
	events chan<- []Todo
}

func (r *Resolver) WatchTodos(ctx context.Context) (<-chan []Todo, error) {
	ch := make(chan []Todo)
	r.Subscriptions.todoSubscribers <- &todosSubscriber{events: ch, stop: ctx.Done()}
	go func(r *Resolver) {
		if todos, err := r.GetTodos(context.Background()); err != nil {
			logrus.WithError(err).Warn("error fetching TODOs")
		} else {
			ch <- todos
		}
	}(r)
	return ch, nil
}

func (r *Resolver) broadcastTodoUpdate() {
	subscribers := map[string]*todosSubscriber{}
	unsubscribe := make(chan string)
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.Subscriptions.todoSubscribers:
			subscribers[ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()] = s
		case e := <-r.Subscriptions.updateTodos:
			for id, s := range subscribers {
				go func(id string, s *todosSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- *e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}
