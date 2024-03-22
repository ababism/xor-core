CREATE TABLE email_notification
(
    uuid           uuid      NOT NULL UNIQUE,
    sender_uuid    uuid      NOT NULL,
    sender_email   varchar   NOT NULL,
    receiver_email varchar   NOT NULL,
    subject        varchar   NOT NULL,
    body           varchar   NOT NULL,
    created_at     timestamp NOT NULL,
    PRIMARY KEY (uuid)
);