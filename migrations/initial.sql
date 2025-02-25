CREATE TABLE IF NOT EXISTS warehouses (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    status INTEGER CHECK (status IN (0, 1))
);

INSERT INTO
    warehouses (id, name, status)
VALUES
    (
        '7d090868-3df5-44e7-9280-3cad6204be59',
        'Primario',
        1
    ),
    (
        'd560d2aa-b9e2-418d-806d-e30784f8d7dc',
        'Secundario',
        1
    );
