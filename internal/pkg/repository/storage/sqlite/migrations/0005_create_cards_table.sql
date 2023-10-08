-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS cards
(
    id  varchar(36) NOT NULL PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id varchar(36) NOT NULL,
    card_number blob NOT NULL,
    card_name blob NOT NULL,
    card_date blob NOT NULL,
    cvc blob NOT NULL,
    comment blob,
    is_deleted  boolean DEFAULT FALSE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS cards;

-- +goose StatementEnd
