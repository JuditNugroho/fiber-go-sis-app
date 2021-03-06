CREATE TABLE IF NOT EXISTS members
(
    id          INT GENERATED BY DEFAULT AS IDENTITY,
    name        VARCHAR(255) NOT NULL,
    phone       VARCHAR(20)  NOT NULL,
    create_time timestamp    NOT NULL DEFAULT NOW(),
    update_time timestamp
);


CREATE INDEX IF NOT EXISTS member_phone_idx ON members (phone);