CREATE TABLE account
(
    id            uuid               NOT NULL,
    login         varchar            NOT NULL UNIQUE,
    password_hash varchar            NOT NULL,
    contacts      jsonb              NOT NULL,
    deleted       bool DEFAULT false NOT NULL,
    PRIMARY KEY (id)
)