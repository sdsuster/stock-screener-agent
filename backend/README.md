# Stock Screener Agent - Microservices Backend

A modern microservices architecture built with **Encore** and **Ent** ORM for stock market data and news analysis.

## Architecture

```
backend/
├── gateway/         # API Gateway - Unified entry point
├── marketdata/      # Market Data Service - Stocks & summaries
├── news/            # News Service - Articles & embeddings
└── docker-compose.yml
```

### Services

1. **Gateway Service** (`/api/*`)
   - Health checks
   - Cross-service queries
   - Unified stock information with news

2. **Market Data Service** (`/marketdata/*`)
   - Stock ticker management
   - Daily stock summaries
   - Historical data queries

3. **News Service** (`/news/*`)
   - News article storage
   - Content chunking for embeddings
   - Semantic search with pgvector

## Technology Stack

- **Encore**: Modern backend framework for Go microservices
- **Ent**: Entity framework with code generation and migrations
- **PostgreSQL**: Primary database with pgvector extension
- **pgvector**: Vector similarity search for news embeddings

## Prerequisites

- Go 1.22+
- Encore CLI: `npm install -g encore`
- Docker & Docker Compose
- PostgreSQL with pgvector

## Quick Start

### 1. Install Encore CLI

```bash
npm install -g encore
```

### 2. Start Database

```bash
cd backend
docker-compose up -d
```

### 3. Generate Ent Code

```bash
# Generate Market Data entities
cd marketdata
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

# Generate News entities
cd ../news
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
```

### 4. Run Migrations

Encore automatically handles database migrations. Create migration files:

```bash
# From backend directory
encore db migrations create --schema marketdata
encore db migrations create --schema news
```

### 5. Run Services

```bash
# From backend directory
encore run
```

The services will be available at:
- Gateway: http://localhost:4000/api/*
- Market Data: http://localhost:4000/marketdata/*
- News: http://localhost:4000/news/*
- Encore Dashboard: http://localhost:9400

## API Endpoints

### Gateway Service

- `GET /health` - Health check for all services
- `GET /api/stocks/:stock_code` - Get comprehensive stock info
- `POST /api/search` - Cross-service search

### Market Data Service

- `GET /marketdata/stocks` - List all stocks
- `GET /marketdata/stocks/:stock_code` - Get single stock
- `POST /marketdata/stocks` - Create/update stock
- `POST /marketdata/summaries` - Get stock summaries with filters

### News Service

- `POST /news` - Create news article
- `GET /news/:id` - Get news article
- `POST /news/list` - List news articles
- `POST /news/search` - Semantic search (vector similarity)
- `GET /news/:news_id/chunks` - Get news chunks

## Ent Migrations

### Generate New Migration

```bash
# After modifying schema files
cd marketdata  # or news
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
```

### Apply Migrations

Encore handles migrations automatically when running:

```bash
encore run
```

Or manually:

```bash
encore db migrate --schema marketdata
encore db migrate --schema news
```

## Database Schema

### Market Data Service

- **stocks** - Stock ticker information
- **stock_summaries** - Daily trading data

### News Service

- **news** - News articles
- **news_chunks** - Content chunks with embeddings (pgvector)

## Development

### Project Structure

```
backend/
├── encore.app           # Encore app configuration
├── go.mod              # Go dependencies
├── gateway/            # API Gateway service
│   └── api.go
├── marketdata/         # Market Data service
│   ├── api.go         # API endpoints
│   ├── service.go     # Service initialization
│   ├── types.go       # Data types
│   └── ent/
│       └── schema/    # Ent schema definitions
│           ├── stock.go
│           └── stocksummary.go
└── news/              # News service
    ├── api.go        # API endpoints
    ├── service.go    # Service initialization
    ├── types.go      # Data types
    └── ent/
        └── schema/   # Ent schema definitions
            ├── news.go
            └── newschunk.go
```

### Adding New Endpoints

1. Define request/response types in `types.go`
2. Add endpoint in `api.go` with `//encore:api` comment
3. Implement business logic

Example:

```go
//encore:api public method=GET path=/news/:id
func GetNews(ctx context.Context, params *GetNewsParams) (*GetNewsResponse, error) {
    // Implementation
}
```

### Modifying Database Schema

1. Update schema files in `ent/schema/`
2. Regenerate Ent code: `go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema`
3. Encore will auto-generate migrations on next run

## Testing

```bash
# Run tests
encore test

# Test specific service
encore test ./marketdata/...
encore test ./news/...
```

## Deployment

```bash
# Deploy to Encore Cloud
encore deploy

# Or export to Docker
encore build docker
```

## Environment Variables

Create `.env` file (optional, Encore handles most config):

```env
# Database is auto-configured by Encore
# Custom settings can be added here
```

## Monitoring

Encore provides built-in monitoring:
- Request traces
- Performance metrics
- Service health
- Database queries

Access at: http://localhost:9400

## Migration from Python

This is a complete rewrite of the Python service (`app/services/news-clawer`) with improvements:

- Type-safe API with Go
- Automatic migrations with Ent
- Better performance
- Built-in tracing and monitoring
- Microservices architecture
- Native vector search support

## Roadmap

- [ ] Implement full Ent repository patterns
- [ ] Add authentication & authorization
- [ ] Implement caching layer (Redis)
- [ ] Add message queue for async tasks
- [ ] Deploy crawler as separate service
- [ ] Add comprehensive tests
- [ ] Set up CI/CD pipeline

## Resources

- [Encore Documentation](https://encore.dev/docs)
- [Ent Documentation](https://entgo.io/docs/getting-started)
- [pgvector Documentation](https://github.com/pgvector/pgvector)
