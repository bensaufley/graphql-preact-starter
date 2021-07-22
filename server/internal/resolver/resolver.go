package resolver

import "context"

type QueryResolver struct {
	Db interface{}
}

func NewRoot(db interface{}) *QueryResolver {
	return &QueryResolver{Db: db}
}

type Bar struct {
	name string
}

func (r QueryResolver) Foo(ctx context.Context) (*Bar, error) {
	return &Bar{}, nil
}

func (b *Bar) Name() string {
	if b == nil {
		return ""
	}
	return b.name
}
