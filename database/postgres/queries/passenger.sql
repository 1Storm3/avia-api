-- name: DeletePassenger :exec
DELETE
FROM passengers
WHERE id = $1;

-- name: UpdatePassenger :one
UPDATE passengers
SET first_name  = COALESCE($1, first_name),
    last_name   = COALESCE($2, last_name),
    middle_name = COALESCE($3, middle_name)
WHERE id = $4 RETURNING *;