CREATE TABLE IF NOT EXISTS stocks (
    id UUID NOT NULL PRIMARY KEY,
    rack_id UUID REFERENCES RACKS(id),
    item_id UUID REFERENCES ITEMS(id),
    quantity INTEGER,
    operation_date TIMESTAMP,
    expiration_date TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);