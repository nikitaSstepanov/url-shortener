-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    alias TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd