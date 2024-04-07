CREATE TABLE bank_accounts
(
    uuid         UUID         NOT NULL,
    account_uuid UUID         NOT NULL,
    login        VARCHAR(255) NOT NULL UNIQUE,
    funds        FLOAT(10)    NOT NULL,
    data         JSONB        NOT NULL,
    status       VARCHAR(255) NOT NULL,
    last_deal_at TIMESTAMP,
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);


CREATE TABLE payments
(
    uuid       UUID        NOT NULL,
    sender     UUID        NOT NULL,
    receiver   UUID        NOT NULL,
    data       JSONB       NOT NULL,
    url        TEXT        NOT NULL,
    status     VARCHAR(32) NOT NULL,
    ended_at   TIMESTAMP   NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid),
    FOREIGN KEY (sender) REFERENCES bank_accounts (uuid),
    FOREIGN KEY (receiver) REFERENCES bank_accounts (uuid)
);

CREATE TABLE products
(
    uuid       UUID,
    name       VARCHAR(255) NOT NULL,
    price      FLOAT(10)    NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);

CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_bank_accounts_updated_at
    BEFORE UPDATE ON bank_accounts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER update_products_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

