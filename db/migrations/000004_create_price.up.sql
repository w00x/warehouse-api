CREATE TABLE IF NOT EXISTS prices (
    id UUID NOT NULL PRIMARY KEY,
    market_id UUID REFERENCES MARKETS(id),
    item_id UUID REFERENCES ITEMS(id),
    price DECIMAL,
    date TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);