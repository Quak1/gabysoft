-- name: CreateFolio :exec
INSERT INTO folios (type, count) 
VALUES (?, 1);

-- name: GetFolios :many
SELECT * FROM folios;

-- name: IncrementCount :exec
UPDATE folios
SET count = count + 1
WHERE id = ?;

