CREATE TABLE IF NOT EXISTS transaction_items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  quantity NUMERIC NOT NULL,
  product_id INTEGER NOT NULL,
  transaction_id INTEGER NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id) ,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);

-- name: AddTransactionItem :exec
INSERT INTO transaction_items (quantity, product_id, transaction_id) 
VALUES (?, ?, ?);
