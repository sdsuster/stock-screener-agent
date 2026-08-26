# Quick Start Guide

## Installation

1. **Install Encore CLI**
   ```bash
   npm install -g encore
   ```

2. **Start Database**
   ```bash
   cd backend
   docker-compose up -d
   ```

3. **Generate Ent Code**
   ```bash
   cd backend
   make generate
   ```

4. **Run Services**
   ```bash
   make dev
   ```

## Services

- **Gateway**: http://localhost:4000/api/*
- **Market Data**: http://localhost:4000/marketdata/*
- **News**: http://localhost:4000/news/*
- **Dashboard**: http://localhost:9400

## Using Make Commands

```bash
make help          # Show all available commands
make install       # Install dependencies
make db-up         # Start database
make generate      # Generate Ent code
make dev           # Run services
make test          # Run tests
make db-down       # Stop database
```

## API Examples

### Get Stock Information
```bash
curl http://localhost:4000/api/stocks/BBCA
```

### Create News
```bash
curl -X POST http://localhost:4000/news \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Market Update",
    "content": "Today market analysis...",
    "url": "https://example.com/news/1"
  }'
```

### Search News
```bash
curl -X POST http://localhost:4000/news/search \
  -H "Content-Type: application/json" \
  -d '{
    "query": "market analysis",
    "limit": 10,
    "similarity": 0.7
  }'
```
