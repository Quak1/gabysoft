-- name: CountProducts :one
SELECT count(*) FROM products;

-- name: CreateProduct :exec
INSERT INTO products (name, code, barcode, description, price) 
VALUES (?, ?, ?, ?, ?);

-- name: GetProducts :many
SELECT * FROM products;
