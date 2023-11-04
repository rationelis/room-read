CREATE TABLE IF NOT EXISTS message (
    id        BIGSERIAL PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL,
    topic     VARCHAR(255) NOT NULL,
    payload   TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);
