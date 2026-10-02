-- name: CreateClient :exec
INSERT INTO clients (name) 
VALUES (?);

-- name: GetClients :many
SELECT * FROM clients;
