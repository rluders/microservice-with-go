-- Create the 'categories' table
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Create an index on the 'deleted_at' column for efficient querying
CREATE INDEX idx_categories_deleted_at ON categories (deleted_at);

-- Create an index on the 'name' column for efficient querying
CREATE INDEX idx_categories_name ON categories (name);