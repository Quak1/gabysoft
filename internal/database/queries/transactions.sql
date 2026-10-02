-- name: CreateTransaction :exec
INSERT INTO transactions (date, client_id, address_id, folio, folio_type) 
VALUES (?, ?, ?, ?, ?);

-- name: GetTransactions :many
SELECT * FROM transactions;

-- name: GetTransactionInfo :one
SELECT * FROM transactions t
JOIN clients c ON t.client_id = c.id
JOIN addresses a ON t.address_id = a.id
JOIN folios f ON t.folio_type = folios.type
WHERE t.id = ?;

-- name: GetTransactionItems :many
SELECT * FROM transactions t
JOIN transaction_items ti ON t.id = ti.transactions_id
JOIN products p ON ti.product_id = p.id
WHERE t.id = ?;
