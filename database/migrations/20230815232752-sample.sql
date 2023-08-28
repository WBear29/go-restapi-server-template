
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE samples (
    id uuid NOT NULL primary key DEFAULT uuid_generate_v4(),
    "name" varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE samples;