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

CREATE TABLE telegram_user
(
    account_uuid uuid    NOT NULL UNIQUE,
    id           bigint  NOT NULL UNIQUE,
    username     varchar NOT NULL UNIQUE,
    first_name   varchar NOT NULL,
    last_name    varchar NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE telegram_chat
(
    id          bigint  NOT NULL UNIQUE,
    title       varchar NOT NULL,
    owner_id    bigint  NOT NULL,
    invite_link varchar,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES telegram_user (id)
);

CREATE TABLE telegram_chat_user
(
    chat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY (chat_id, user_id),
    FOREIGN KEY (chat_id) REFERENCES telegram_chat (id),
    FOREIGN KEY (user_id) REFERENCES telegram_user (id)
);
