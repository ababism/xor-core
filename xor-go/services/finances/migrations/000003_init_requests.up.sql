CREATE TABLE payout_requests
(
    uuid       UUID      NOT NULL DEFAULT gen_random_uuid(),
    receiver   UUID      NOT NULL,
    amount     FLOAT(10) NOT NULL,
    data       JSONB     NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);

CREATE TABLE purchase_requests
(
    uuid        UUID      NOT NULL DEFAULT gen_random_uuid(),
    sender      UUID,
    receiver    UUID,
    status      UUID      NOT NULL,
    amount      FLOAT(10) NOT NULL,
    webhook_url TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);

CREATE TABLE purchase_requests_products
(
    request_uuid UUID,
    product_uuid UUID,
    PRIMARY KEY (request_uuid, product_uuid),
    FOREIGN KEY (request_uuid) REFERENCES purchase_requests (uuid),
    FOREIGN KEY (product_uuid) REFERENCES products (uuid)
);
