CREATE TABLE feedback_resource
(
    uuid            uuid      NOT NULL UNIQUE,
    name            varchar   NOT NULL,
    description     varchar   NOT NULL,
    created_by_uuid uuid      NOT NULL,
    created_at      timestamp NOT NULL,
    active          bool      NOT NULL,
    PRIMARY KEY (uuid)
);

CREATE TABLE feedback_item
(
    uuid            uuid      NOT NULL UNIQUE,
    resource_uuid   uuid      NOT NULL,
    created_by_uuid uuid      NOT NULL,
    created_at      timestamp NOT NULL,
    text            varchar   NOT NULL,
    rating          int       NOT NULL,
    active          bool      NOT NULL,
    PRIMARY KEY (uuid),
    FOREIGN KEY (resource_uuid) REFERENCES feedback_resource (uuid)
);