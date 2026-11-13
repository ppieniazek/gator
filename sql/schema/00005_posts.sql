-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    title TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    description TEXT,
    published_at TIMESTAMPTZ,
    feed_id UUID NOT NULL REFERENCES feeds (id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS posts;
