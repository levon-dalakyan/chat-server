-- +goose Up
CREATE TABLE chats (
    id uuid PRIMARY KEY,
    usersnames text[],
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);

-- +goose Down
DROP TABLE chats;

