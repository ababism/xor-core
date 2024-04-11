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

CREATE TABLE role
(
    uuid            uuid    NOT NULL UNIQUE,
    name            varchar NOT NULL UNIQUE,
    created_by_uuid uuid    NOT NULL,
    created_at      time    NOT NULL,
    active          bool    NOT NULL,
    PRIMARY KEY (uuid),
    FOREIGN KEY (created_by_uuid) REFERENCES account (uuid)
);

CREATE TABLE account_role
(
    account_uuid uuid NOT NULL,
    role_uuid    uuid NOT NULL,
    UNIQUE (account_uuid, role_uuid),
    PRIMARY KEY (account_uuid, role_uuid),
    FOREIGN KEY (account_uuid) REFERENCES account (uuid),
    FOREIGN KEY (role_uuid) REFERENCES role (uuid)
);
