CREATE TABLE IF NOT EXISTS history (
    id BIGSERIAL PRIMARY KEY,
    patent_id BIGSERIAL NOT NULL REFERENCES patients ON DELETE CASCADE,
    notes TEXT NOT NULL,
    treatment_id BIGSERIAL NOT NULL REFERENCES treatments ON DELETE CASCADE,
    admitted_date timestamp(0) with time zone NOT NULL DEFAULT NOW() NOT NULL,
    discharged_date timestamp(0) with time zone NOT NULL DEFAULT NOW() NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);