-- +goose Up
CREATE TABLE users (
    id uuid,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    user_name text NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users;
