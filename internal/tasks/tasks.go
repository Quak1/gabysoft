package tasks

import (
	"context"
	"gabysoft/internal/store"
)

type Tasks struct {
	Products *Product
}

func NewTasks() *Tasks {
	return &Tasks{
		Products: &Product{},
	}
}

func (t *Tasks) Inject(ctx context.Context, q *store.Queries) {
	t.Products.ctx = ctx
	t.Products.query = q
}
