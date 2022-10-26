CREATE TABLE IF NOT EXISTS markets (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);