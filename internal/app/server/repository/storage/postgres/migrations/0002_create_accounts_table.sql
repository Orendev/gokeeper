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
    login         varchar(50) NOT NULL,
    password      varchar NOT NULL,
    comment       varchar(255),
    url   text,
    is_deleted  boolean      DEFAULT FALSE                  NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS accounts;

-- +goose StatementEnd
