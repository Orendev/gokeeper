-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users"
(
    id         varchar(36) NOT NULL PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    email      varchar(255) NOT NULL,
    role       varchar(50) NOT NULL,
    password   varchar NOT NULL,
    name       varchar(50) NOT NULL,
    token      varchar,
    CONSTRAINT users_email_unique UNIQUE (email)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "users";

-- +goose StatementEnd
