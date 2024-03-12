CREATE TABLE payout_requests
(
    uuid     UUID PRIMARY KEY,
    receiver UUID           NOT NULL,
    amount   NUMERIC(15, 2) NOT NULL,
    started_at TIMESTAMP      NOT NULL,
    data     JSONB          NOT NULL
);

CREATE TABLE purchase_requests
(
    uuid        UUID PRIMARY KEY,
    sender      UUID      NOT NULL,
    receiver    UUID      NOT NULL,
    webhook_url TEXT      NOT NULL,
    started_at    TIMESTAMP NOT NULL
);
