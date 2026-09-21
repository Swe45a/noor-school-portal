-- id_number and employee_number are the original identifiers imported from Excel and must
-- never be regenerated or reassigned; id is an internal surrogate key only. A teacher must be
-- authenticated with both values together, so the pair is what the login lookup matches on.
CREATE TABLE teachers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_number VARCHAR(50) UNIQUE NOT NULL,
    employee_number VARCHAR(50) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    account_username VARCHAR(100) UNIQUE NOT NULL,
    account_password_hash VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_teachers_id_number ON teachers(id_number);
CREATE INDEX idx_teachers_employee_number ON teachers(employee_number);
