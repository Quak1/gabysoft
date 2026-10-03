package tasks

import (
	"context"
	"gabysoft/internal/store"
)

type Tasks struct {
	Products    *Product
	Address     *Address
	Client      *Client
	Folio       *Folio
	Transaction *Transaction
}

func NewTasks() *Tasks {
	return &Tasks{
		Products:    &Product{},
		Address:     &Address{},
		Client:      &Client{},
		Folio:       &Folio{},
		Transaction: &Transaction{},
	}
}

func (t *Tasks) Inject(ctx context.Context, q *store.Queries) {
	t.Products.ctx = ctx
	t.Products.query = q

	t.Address.ctx = ctx
	t.Address.query = q

	t.Client.ctx = ctx
	t.Client.query = q

	t.Folio.ctx = ctx
	t.Folio.query = q

	t.Transaction.ctx = ctx
	t.Transaction.query = q
}
