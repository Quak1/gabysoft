package tasks

import (
	"context"
	"gabysoft/internal/store"
)

type Address struct {
	ctx   context.Context
	query *store.Queries
}

func (a *Address) Create(address store.CreateAddressParams) error {
	return a.query.CreateAddress(a.ctx, address)
}

func (a *Address) GetAll() ([]store.Address, error) {
	return a.query.GetAddressess(a.ctx)
}
