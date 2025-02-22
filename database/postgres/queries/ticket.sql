-- name: GetAllTickets :many
SELECT id,
       passenger_id,
       order_number,
       departure,
       destination,
       service_provider,
       departure_date,
       arrival_date,
       order_date
FROM tickets
ORDER BY order_date LIMIT $1
OFFSET $2;

-- name: GetCountTickets :one
SELECT count(*) AS count
FROM tickets;

-- name: UpdateTicket :one
UPDATE tickets
SET order_number     = COALESCE($1, order_number),
    departure        = COALESCE($2, departure),
    destination      = COALESCE($3, destination),
    service_provider = COALESCE($4, service_provider),
    departure_date   = COALESCE($5, departure_date),
    arrival_date     = COALESCE($6, arrival_date),
    order_date       = COALESCE($7, order_date)
WHERE id = $8 RETURNING *;

-- name: GetAllByPassenger :many
SELECT id,
       passenger_id,
       order_number,
       departure,
       destination,
       service_provider,
       departure_date,
       arrival_date,
       order_date
FROM tickets
WHERE passenger_id = $1;

-- name: DeleteTicket :exec
DELETE
FROM tickets
WHERE id = $1;

-- name: GetOneTicket :one
SELECT t.*,
       sqlc.embed(p),
       sqlc.embed(d)
FROM tickets t
         left JOIN passengers p ON p.id = t.passenger_id
         left JOIN documents d ON d.passenger_id = p.id
WHERE t.id = $1;

-- name: GetPassengerReport :many
SELECT t.order_date,
       t.departure_date,
       t.order_number,
       t.departure,
       t.destination,
       CASE
           WHEN t.departure_date <= NOW() THEN true
           ELSE false
           END AS service_provided
FROM tickets t
         JOIN passengers p ON t.passenger_id = p.id
WHERE p.id = $1
  AND (
    (t.order_date < $2 AND t.departure_date BETWEEN $2 AND $3)
        OR (t.order_date BETWEEN $2 AND $3)
    )
ORDER BY t.order_date;
