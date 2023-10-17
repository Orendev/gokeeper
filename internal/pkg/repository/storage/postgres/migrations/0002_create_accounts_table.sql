-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS accounts
(
    id           uuid DEFAULT gen_random_uuid() NOT NULL
        CONSTRAINT pk_account
            PRIMARY KEY,
    created_at   timestamp,
    updated_at  timestamp,
    user_id  uuid                                not null
        constraint fk_user_id
            references users
            on delete cascade,
    title   varchar(100) DEFAULT '',
    login         bytea NOT NULL,
    password      bytea NOT NULL,
    comment       bytea,
    url   bytea,
    is_deleted  boolean      DEFAULT FALSE                  NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS accounts;

-- +goose StatementEnd
