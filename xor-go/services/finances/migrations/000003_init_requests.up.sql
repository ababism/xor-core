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
    sender      UUID,
    receiver    UUID,
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
