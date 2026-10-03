package tasks

import (
	"context"
	"gabysoft/internal/store"
)

type Client struct {
	ctx   context.Context
	query *store.Queries
}

func (c *Client) Create(name string) error {
	return c.query.CreateClient(c.ctx, name)
}

func (c *Client) GetAll() ([]store.Client, error) {
	return c.query.GetClients(c.ctx)
}
