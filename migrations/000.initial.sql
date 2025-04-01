CREATE TABLE IF NOT EXISTS warehouses (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    status INTEGER CHECK (status IN (0, 1)),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
)