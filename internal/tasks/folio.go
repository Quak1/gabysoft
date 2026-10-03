package tasks

import (
	"context"
	"gabysoft/internal/store"
)

type Folio struct {
	ctx   context.Context
	query *store.Queries
}

func (f *Folio) Create(folio string) error {
	return f.query.CreateFolio(f.ctx, folio)
}

func (f *Folio) GetAll() ([]store.Folio, error) {
	return f.query.GetFolios(f.ctx)
}

func (f *Folio) Increment(folio string) error {
	return f.query.IncrementFolioCount(f.ctx, folio)
}
