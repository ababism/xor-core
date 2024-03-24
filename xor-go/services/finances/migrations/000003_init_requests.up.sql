CREATE TABLE payout_requests
(
    uuid       UUID,
    receiver   UUID      NOT NULL,
    amount     FLOAT(10) NOT NULL,
    data       JSONB     NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);

CREATE TABLE purchase_requests
(
    uuid        UUID,
    sender      UUID      NOT NULL,
    receiver    UUID      NOT NULL,
    webhook_url TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
