CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE accounts (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name         TEXT NOT NULL,
    account_no   TEXT NOT NULL UNIQUE,
    ifsc_code    TEXT NOT NULL,
    pin_code     TEXT NOT NULL,
    balance      NUMERIC(18, 2) NOT NULL DEFAULT 0,
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at   TIMESTAMP WITH TIME ZONE     -- NULL means active, NOT NULL means soft deleted
);
