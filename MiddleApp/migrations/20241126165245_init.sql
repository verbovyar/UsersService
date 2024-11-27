-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    id serial PRIMARY KEY,
    name text NOT NULL,
    surname text NOT NULL,
    age integer NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
-- +goose StatementEnd
