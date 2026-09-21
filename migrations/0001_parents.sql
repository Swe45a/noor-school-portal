CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- parent_id is the original identifier imported from Excel and must never be regenerated or
-- reassigned; id is an internal surrogate key only.
CREATE TABLE parents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    parent_id VARCHAR(50) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    account_username VARCHAR(100) UNIQUE NOT NULL,
    account_password_hash VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_parents_parent_id ON parents(parent_id);
