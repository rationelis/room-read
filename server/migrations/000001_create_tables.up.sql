CREATE TABLE IF NOT EXISTS message (
    id        INTEGER NOT NULL PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL,
    topic     VARCHAR(255) NOT NULL,
    payload   BLOB NOT NULL,
    timestamp TIMESTAMP NOT NULL
);
