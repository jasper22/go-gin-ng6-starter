CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    created_at bigint,
    updated_at bigint,
    first_name text,
    last_name text,
    username text NOT NULL UNIQUE,
    password text,
    role integer,
    xxx_unrecognized bytea,
    xxx_sizecache integer
);

