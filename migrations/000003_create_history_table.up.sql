CREATE TABLE IF NOT EXISTS history (
    id BIGSERIAL PRIMARY KEY,
    patent_id BIGSERIAL NOT NULL REFERENCES patients ON DELETE CASCADE,
    notes TEXT NOT NULL,
    treatment_id BIGSERIAL NOT NULL REFERENCES treatments ON DELETE CASCADE,
    admitted_date DATETIME NOT NULL,
    discharged_date DATETIME NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);