CREATE TABLE IF NOT EXISTS racks (
   id UUID NOT NULL PRIMARY KEY,
   name VARCHAR,
   code VARCHAR,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);