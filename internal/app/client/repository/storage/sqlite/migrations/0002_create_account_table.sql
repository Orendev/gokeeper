-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "accounts"
(
    id  varchar(36) NOT NULL PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    title   varchar(100),
    login   varchar(50) NOT NULL,
    password      varchar NOT NULL,
    comment       varchar(255),
    url   text,
    is_deleted  boolean DEFAULT FALSE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "accounts";

-- +goose StatementEnd
