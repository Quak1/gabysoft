package tasks

import (
	"context"
	"gabysoft/internal/store"
)

type Product struct {
	ctx   context.Context
	query *store.Queries
}

func (p *Product) GetAll() ([]store.Product, error) {
	return p.query.GetProducts(p.ctx)
}

func (p *Product) Create(newP store.CreateProductParams) error {
	return p.query.CreateProduct(p.ctx, store.CreateProductParams{
		Name:        newP.Name,
		Code:        newP.Code,
		Barcode:     newP.Barcode,
		Description: newP.Description,
		Price:       newP.Price,
	})
}

func (p *Product) Update(newP store.CreateProductParams, id int64) error {
	return p.query.UpdateProduct(p.ctx, store.UpdateProductParams{
		Name:        newP.Name,
		Code:        newP.Code,
		Barcode:     newP.Barcode,
		Description: newP.Description,
		Price:       newP.Price,
		ID:          id,
	})
}

func (p *Product) Delete(id int64) error {
	return p.query.DeleteProduct(p.ctx, id)
}
