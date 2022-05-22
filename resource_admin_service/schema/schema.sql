CREATE TABLE IF NOT EXISTS resources(
    id TEXT UNIQUE NOT NULL,
    title TEXT,
    author TEXT,
    content TEXT,
    created_at TIMESTAMP,
    tags []TEXT
);