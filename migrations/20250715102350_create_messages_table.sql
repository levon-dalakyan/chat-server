-- +goose Up
CREATE TABLE messages (
    id bigint PRIMARY KEY,
    chat_id bigint NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    "from" text NOT NULL,
    text text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE messages;

