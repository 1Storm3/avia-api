-- name: DeleteDocument :exec
DELETE
FROM documents
WHERE id = $1;

-- name: GetDocumentsByPassenger :many
SELECT *
FROM documents
WHERE passenger_id = $1;

-- name: UpdateDocument :one
UPDATE documents
SET type = COALESCE($1, type),
    number = COALESCE($2, number)
WHERE id = $3 RETURNING *;