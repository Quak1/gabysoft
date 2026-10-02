-- name: CreateAddress :exec
INSERT INTO addresses (city, state, country, postal_code, street, stree_number, phone_number, tax_id, email) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAddressess :many
SELECT * FROM addresses;
