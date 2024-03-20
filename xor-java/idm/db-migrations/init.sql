CREATE TABLE account
(
    uuid              uuid    NOT NULL UNIQUE,
    email             varchar NOT NULL UNIQUE,
    password_hash     varchar NOT NULL,
    active            bool    NOT NULL,
    first_name        varchar,
    last_name         varchar,
    telegram_username varchar,
    PRIMARY KEY (uuid)
);