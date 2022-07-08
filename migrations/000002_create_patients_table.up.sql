CREATE TABLE IF NOT EXISTS patients (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    trainer_id BIGSERIAL NOT NULL REFERENCES trainers ON DELETE CASCADE,
    pokemon_id BIGINT NOT NULL,
    is_checked_in BOOLEAN NOT NULL DEFAULT false,
    ward TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
);