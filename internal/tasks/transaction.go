package tasks

import (
	"context"
	"gabysoft/internal/store"
	"time"
)

type Transaction struct {
	ctx   context.Context
	query *store.Queries
}

func (t *Transaction) Create(transaction store.CreateTransactionParams) error {
	return t.query.CreateTransaction(t.ctx, transaction)
}

func (t *Transaction) GetAll() ([]store.Transaction, error) {
	return t.query.GetTransactions(t.ctx)
}

func (t *Transaction) AddItem(item store.AddTransactionItemParams) error {
	return t.query.AddTransactionItem(t.ctx, item)
}

type TransactionInfo struct {
	ID        int64
	Date      time.Time
	Folio     int64
	FolioType string
	Client    store.Client
	Address   store.Address
}

type TransactionItem struct {
	store.Product
	Quantity float64
}

type FullTransaction struct {
	Info  TransactionInfo
	Items []TransactionItem
}

func (t *Transaction) Get(id int64) (*FullTransaction, error) {
	info, err := t.query.GetTransactionInfo(t.ctx, id)
	if err != nil {
		return nil, err
	}

	dbItems, err := t.query.GetTransactionItems(t.ctx, id)
	if err != nil {
		return nil, err
	}

	items := make([]TransactionItem, len(dbItems))
	for i, item := range dbItems {
		items[i] = TransactionItem{
			Product: store.Product{
				ID:          item.ID,
				Name:        item.Name,
				Code:        item.Code,
				Barcode:     item.Barcode,
				Description: item.Description,
				Price:       item.Price,
			},
			Quantity: item.Quantity,
		}
	}

	transaction := &FullTransaction{
		Info: TransactionInfo{
			ID:        info.ID,
			Date:      info.Date,
			Folio:     info.Folio,
			FolioType: info.FolioType,
			Client: store.Client{
				ID:   info.ClientID,
				Name: info.Name,
			},
			Address: store.Address{
				ID:          info.AddressID,
				City:        info.City,
				State:       info.State,
				Country:     info.Country,
				PostalCode:  info.PostalCode,
				Street:      info.Street,
				StreeNumber: info.StreeNumber,
				PhoneNumber: info.PhoneNumber,
				TaxID:       info.TaxID,
				Email:       info.Email,
			},
		},
		Items: items,
	}

	return transaction, nil
}
