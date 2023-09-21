-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA IF NOT EXISTS keeper;

CREATE TABLE IF NOT EXISTS keeper.user
(
    id            uuid      DEFAULT gen_random_uuid() NOT NULL
        CONSTRAINT pk_user
            PRIMARY KEY,
    created_at   timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp DEFAULT CURRENT_TIMESTAMP,
    email        varchar(250) NOT NULL
        constraint users_email_unique
            unique,
    role        varchar(50) NOT NULL,
    password     varchar(50) NOT NULL,
    name         varchar(100) DEFAULT '':: character varying NOT NULL,
    surname      varchar(250) DEFAULT '':: character varying NOT NULL,
    patronymic   varchar(100) DEFAULT '':: character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS keeper.account
(
    id           uuid         DEFAULT gen_random_uuid()      NOT NULL
        CONSTRAINT pk_account
            PRIMARY KEY,
    created_at   timestamp,
    updated_at  timestamp,
    user_id  uuid                                not null
        constraint fk_user_id
            references keeper.user
            on delete cascade,
    title   varchar(100) DEFAULT '':: character varying NOT NULL,
    login         varchar(50)  DEFAULT '':: character varying NOT NULL,
    password      varchar(100) DEFAULT '':: character varying NOT NULL,
    comment       varchar(250),
    web_address    text,
    version        int
        CONSTRAINT version_check
            CHECK (version >= 0),
    is_deleted  boolean      DEFAULT FALSE                  NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS keeper.account;

DROP TABLE IF EXISTS keeper.user;

DROP SCHEMA IF EXISTS keeper;

-- +goose StatementEnd
