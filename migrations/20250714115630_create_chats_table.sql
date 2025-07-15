-- +goose Up
CREATE TABLE chats (
    id bigint PRIMARY KEY,
    usernames text[],
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);

-- +goose Down
DROP TABLE chats;

