CREATE TABLE discounts
(
    uuid        UUID        NOT NULL,
    created_by  UUID        NOT NULL,
    percent     FLOAT(2)    NOT NULL,
    stand_alone BOOLEAN     NOT NULL,
    started_at  TIMESTAMP   NOT NULL,
    ended_at    TIMESTAMP   NOT NULL,
    status      VARCHAR(32) NOT NULL,
    PRIMARY KEY (uuid)
);

CREATE TABLE discounts_products
(
    discount_uuid UUID,
    product_uuid  UUID,
    PRIMARY KEY (discount_uuid, product_uuid),
    FOREIGN KEY (discount_uuid) REFERENCES discounts (uuid),
    FOREIGN KEY (product_uuid) REFERENCES products (uuid)
);
