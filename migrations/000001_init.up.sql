CREATE TABLE IF NOT EXISTS user
(
    id              integer PRIMARY KEY,
    created_at      text          NOT NULL,
    updated_at      text          NOT NULL,
    phone_number    text          NOT NULL,
    email           text          NOT NULL,
    hashed_password character(60) NOT NULL,
    version         integer       NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS stock
(
    id         integer PRIMARY KEY,
    created_at text    NOT NULL,
    updated_at text    NOT NULL,
    symbol     text    NOT NULL,
    company    text    NOT NULL,
    version    integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS stock_price
(
    id         integer PRIMARY KEY,
    created_at text    NOT NULL,
    updated_at text    NOT NULL,
    stock_id   bigint REFERENCES stock (id),
    date       text    NOT NULL,
    open       integer NOT NULL,
    high       integer NOT NULL,
    low        integer NOT NULL,
    close      integer NOT NULL,
    version    integer NOT NULL DEFAULT 1
);
