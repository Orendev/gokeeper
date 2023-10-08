-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "accounts"
(
    id  varchar(36) NOT NULL PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    title   varchar(100),
    user_id varchar(36) NOT NULL,
    login   blob NOT NULL,
    password      blob NOT NULL,
    comment       blob,
    url   blob,
    is_deleted  boolean DEFAULT FALSE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "accounts";

-- +goose StatementEnd
