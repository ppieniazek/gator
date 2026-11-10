-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users (id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS feeds;
