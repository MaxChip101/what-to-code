DROP TABLE IF EXISTS ideas CASCADE;

CREATE TABLE ideas (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    tags VARCHAR(50)[] 
);

CREATE INDEX idx_ideas_tags ON ideas USING GIN (tags);
