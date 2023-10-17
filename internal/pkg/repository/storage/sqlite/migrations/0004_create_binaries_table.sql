-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "binaries"
(
    id  varchar(36) NOT NULL PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id varchar(36) NOT NULL,
    title   varchar(100),
    data    blob NOT NULL,
    comment blob,
    is_deleted  boolean DEFAULT FALSE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "binaries";

-- +goose StatementEnd
