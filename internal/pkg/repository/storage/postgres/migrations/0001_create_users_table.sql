-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS users
(
    id            uuid      DEFAULT gen_random_uuid() NOT NULL CONSTRAINT pk_user PRIMARY KEY,
    password   varchar NOT NULL,
    email        varchar(250) NOT NULL
        constraint users_email_unique unique,
    name       varchar(100) DEFAULT '':: character varying NOT NULL,
    role       varchar(50) NOT NULL,
    token      varchar,
    created_at   timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS users;


-- +goose StatementEnd
