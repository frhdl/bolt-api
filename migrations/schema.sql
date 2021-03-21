CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    name Text,
    email VARCHAR(256),
    client_id VARCHAR(16),
    client_secret VARCHAR(16),
    create_at TIMESTAMP,
    update_at TIMESTAMP,
    CONSTRAINT user_unique UNIQUE (client_id, email)
);
