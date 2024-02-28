CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    hashed_password bytea NOT NULL,
    created timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);