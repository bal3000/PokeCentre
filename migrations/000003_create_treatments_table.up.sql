CREATE TABLE IF NOT EXISTS treatments (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    notes TEXT NOT NULL,
    effective_types TEXT[] NOT NULL,
    avoid_types TEXT[] NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);