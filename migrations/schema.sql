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

CREATE TABLE IF NOT EXISTS Projects (
    id SERIAL PRIMARY KEY,
    name TEXT,
    user_id BIGINT,
    create_at TIMESTAMP,
    CONSTRAINT project_unique UNIQUE(name, user_id),
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Tasks (
    id SERIAL PRIMARY KEY,
    description TEXT,
    user_id BIGINT,
    project_id BIGINT,
    create_at TIMESTAMP,
    finish_at TIMESTAMP,
    done boolean,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES Projects(id) ON DELETE CASCADE
);
