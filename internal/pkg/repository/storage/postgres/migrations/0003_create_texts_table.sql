-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS texts
(
    id           uuid DEFAULT gen_random_uuid() NOT NULL
        CONSTRAINT pk_text
            PRIMARY KEY,
    created_at   timestamp,
    updated_at  timestamp,
    user_id  uuid                                not null
        constraint fk_user_id
            references users
            on delete cascade,
    title   varchar(100) DEFAULT '',
    data    bytea NOT NULL,
    comment bytea,
    is_deleted  boolean      DEFAULT FALSE                  NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS texts;

-- +goose StatementEnd
