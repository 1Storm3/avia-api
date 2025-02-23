package pgx

const GetOneTicket = `
        SELECT
            t.id, t.order_number, t.destination, t.departure,
t.service_provider, t.departure_date, t.arrival_date, t.order_date,
            p.id, p.first_name, p.last_name, p.middle_name,
            d.id, d.passenger_id, d.type, d.number
        FROM tickets t
        LEFT JOIN passengers p ON p.id = t.passenger_id
        LEFT JOIN documents d ON d.passenger_id = p.id
        WHERE t.id = $1
    `
