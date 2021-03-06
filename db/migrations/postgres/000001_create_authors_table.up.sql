CREATE TABLE IF NOT EXISTS authors(
  id          UUID PRIMARY KEY,
  name        VARCHAR(255) NOT NULL,
  email       VARCHAR(255) UNIQUE NOT NULL,
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP DEFAULT NULL,
  deleted_at  TIMESTAMP DEFAULT NULL
);