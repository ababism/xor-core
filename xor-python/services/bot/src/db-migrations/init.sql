CREATE TABLE telegram_user
(
--     account_uuid uuid    NOT NULL UNIQUE,
    id       bigint  NOT NULL UNIQUE,
    username varchar NOT NULL UNIQUE,
--     first_name   varchar NOT NULL,
--     last_name    varchar NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE telegram_chat
(
    id          bigint NOT NULL UNIQUE,
    title       varchar,
--     owner_id    bigint  NOT NULL,
    invite_link varchar,
    PRIMARY KEY (id)
--     FOREIGN KEY (owner_id) REFERENCES telegram_user (id)
);

CREATE TABLE telegram_chat_user
(
    chat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY (chat_id, user_id),
    FOREIGN KEY (chat_id) REFERENCES telegram_chat (id),
    FOREIGN KEY (user_id) REFERENCES telegram_user (id)
);
