-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS accounts
(
    id           uuid         DEFAULT gen_random_uuid()      NOT NULL
        CONSTRAINT pk_account
            PRIMARY KEY,
    created_at   timestamp,
    updated_at  timestamp,
    user_id  uuid                                not null
        constraint fk_user_id
            references users
            on delete cascade,
    title   varchar(100) DEFAULT '':: character varying NOT NULL,
    login         varchar(50)  DEFAULT '':: character varying NOT NULL,
    password      varchar DEFAULT '':: character varying NOT NULL,
    comment       varchar(250),
    web_address   text,
    version        int
        CONSTRAINT version_check
            CHECK (version >= 0),
    is_deleted  boolean      DEFAULT FALSE                  NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS accounts;

-- +goose StatementEnd
