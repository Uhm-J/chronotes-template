-- Database initialization script for Chronotes
-- This script runs when the PostgreSQL Docker container starts

-- Create the main database (already created by POSTGRES_DB)
-- CREATE DATABASE chronotes;

-- Connect to the database
\c chronotes;

-- Create extensions if needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- The application will handle table creation via migrations
-- But we can add any additional setup here

-- Create a read-only user for reporting (optional)
-- CREATE USER chronotes_readonly WITH PASSWORD 'readonly_pass';
-- GRANT CONNECT ON DATABASE chronotes TO chronotes_readonly;
-- GRANT USAGE ON SCHEMA public TO chronotes_readonly;
-- GRANT SELECT ON ALL TABLES IN SCHEMA public TO chronotes_readonly;
-- ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO chronotes_readonly;

-- Print success message
\echo 'Database initialization completed successfully!' 