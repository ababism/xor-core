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
