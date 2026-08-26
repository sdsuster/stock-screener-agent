-- Enable pgvector extension
CREATE EXTENSION IF NOT EXISTS vector;

-- Create separate databases for each service
CREATE DATABASE marketdata;
CREATE DATABASE news;

-- Connect to marketdata database
\c marketdata;
CREATE EXTENSION IF NOT EXISTS vector;

-- Connect to news database
\c news;
CREATE EXTENSION IF NOT EXISTS vector;

-- Grant permissions
GRANT ALL PRIVILEGES ON DATABASE marketdata TO postgres;
GRANT ALL PRIVILEGES ON DATABASE news TO postgres;
