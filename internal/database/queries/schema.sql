CREATE TABLE IF NOT EXISTS products (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  code TEXT NOT NULL,
  barcode TEXT NOT NULL,
  description TEXT NOT NULL,
  price NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS clients (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS addresses (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  city TEXT NOT NULL,
  state TEXT NOT NULL,
  country TEXT NOT NULL,
  postal_code TEXT NOT NULL,
  street TEXT NOT NULL,
  stree_number TEXT NOT NULL,
  phone_number TEXT NOT NULL,
  tax_id TEXT NOT NULL,
  email TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  date DATE NOT NULL,
  folio INTEGER NOT NULL,
  client_id INTEGER NOT NULL,
  address_id INTEGER NOT NULL,
  folio_type TEXT NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients(id),
  FOREIGN KEY (address_id) REFERENCES addresses(id),
  FOREIGN KEY (folio_type) REFERENCES folios(type) 
);

CREATE TABLE IF NOT EXISTS transaction_items (
  quantity NUMERIC NOT NULL,
  product_id INTEGER NOT NULL,
  transaction_id INTEGER NOT NULL,
  PRIMARY KEY (product_id, transaction_id),
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);

CREATE TABLE IF NOT EXISTS folios (
  type TEXT PRIMARY KEY,
  count INTEGER NOT NULL
);
