CREATE TABLE IF NOT EXISTS trainers (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  address TEXT NOT NULL,
  nhs_number TEXT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME
);