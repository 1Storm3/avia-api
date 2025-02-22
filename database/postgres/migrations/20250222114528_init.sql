-- +goose Up
CREATE TABLE passengers
(
    id          UUID PRIMARY KEY,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255) NOT NULL
);

CREATE TABLE tickets
(
    id               UUID PRIMARY KEY,
    passenger_id     UUID                     NOT NULL REFERENCES passengers (id) ON DELETE CASCADE,
    order_number     VARCHAR(255)             NOT NULL,
    departure        VARCHAR(255)             NOT NULL,
    destination      VARCHAR(255)             NOT NULL,
    service_provider VARCHAR(255)             NOT NULL,
    departure_date   timestamp with time zone NOT NULL,
    arrival_date     timestamp with time zone NOT NULL,
    order_date       timestamp with time zone NOT NULL,
    CONSTRAINT fk_passenger FOREIGN KEY (passenger_id) REFERENCES passengers (id) ON DELETE CASCADE
);

CREATE TABLE documents
(
    id           UUID PRIMARY KEY,
    passenger_id UUID NULL REFERENCES passengers (id) ON DELETE CASCADE,
    type         VARCHAR(255) NOT NULL,
    number       VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE documents;
DROP TABLE tickets;
DROP TABLE passengers;
