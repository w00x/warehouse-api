CREATE TABLE IF NOT EXISTS items (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR,
    unit_size_presentation VARCHAR,
    size_presentation INTEGER,
    code VARCHAR,
    container VARCHAR,
    photo VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);