-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"
(
    id         varchar(36) NOT NULL PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    email      varchar(255) NOT NULL,
    role       varchar(50) NOT NULL,
    password   varchar(50) NOT NULL,
    name       varchar(50) NOT NULL,
    token      varcher(255),
    CONSTRAINT users_email_unique UNIQUE (email)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "user";

-- +goose StatementEnd
