CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT users_email_unique UNIQUE (email),
    CONSTRAINT users_email_not_empty CHECK (length(trim(email)) > 0),
    CONSTRAINT users_full_name_not_empty CHECK (length(trim(full_name)) > 0),
    CONSTRAINT users_password_hash_not_empty CHECK (length(trim(password_hash)) > 0)
);

CREATE INDEX users_email_idx ON users (email);
