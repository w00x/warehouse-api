CREATE TABLE IF NOT EXISTS inventories (
     id UUID NOT NULL PRIMARY KEY,
     operation_date TIMESTAMP,
     created_at TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP
);