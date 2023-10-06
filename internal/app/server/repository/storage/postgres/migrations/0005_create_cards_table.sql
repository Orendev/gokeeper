-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS cards
(
    id           uuid DEFAULT gen_random_uuid() NOT NULL
        CONSTRAINT pk_card
            PRIMARY KEY,
    created_at   timestamp,
    updated_at  timestamp,
    user_id  uuid                                not null
        constraint fk_user_id
            references users
            on delete cascade,
    card_number BYTEA NOT NULL,
    card_name BYTEA NOT NULL,
    card_date BYTEA NOT NULL,
    cvc BYTEA NOT NULL,
    comment bytea,
    is_deleted  boolean      DEFAULT FALSE                  NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS cards;

-- +goose StatementEnd
