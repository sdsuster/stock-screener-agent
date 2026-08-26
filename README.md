# Stock Screener Agent

An intelligent stock screening and analysis system with microservices architecture.

## Project Structure

```
.
├── app/                    # Python application (legacy)
│   ├── main.py
│   ├── agents/            # AI agents
│   ├── config/            # Configuration
│   ├── services/          # Legacy services
│   └── workers/           # Background workers
│
└── backend/               # Go microservices (Encore + Ent)
    ├── gateway/           # API Gateway
    ├── marketdata/        # Market Data Service
    ├── news/              # News Service
    └── docker-compose.yml
```

## Backend Services (New)

The new microservices architecture is built with:
- **Encore**: Modern Go backend framework
- **Ent**: ORM with code generation and migrations
- **PostgreSQL + pgvector**: Vector database for semantic search

### Quick Start

See [backend/README.md](backend/README.md) for detailed documentation.

```bash
cd backend
make install      # Install dependencies
make db-up        # Start database
make generate     # Generate Ent code
make dev          # Run services
```

### Services

1. **Gateway** - Unified API entry point
2. **Market Data** - Stock tickers and daily summaries
3. **News** - News articles with semantic search

## Features

- ✅ Microservices architecture
- ✅ Type-safe APIs with Go
- ✅ Automatic database migrations (Ent)
- ✅ Vector similarity search (pgvector)
- ✅ Built-in monitoring and tracing
- ✅ Docker support
- 🚧 News crawler integration
- 🚧 AI-powered stock screening

## Development

### Backend (Go Microservices)
```bash
cd backend
make dev          # Run in development mode
make test         # Run tests
```

### Python App (Legacy)
```bash
cd app
pip install -r ../requirements.txt
python main.py
```

## Documentation

- [Backend README](backend/README.md) - Complete backend documentation
- [Backend Quick Start](backend/QUICKSTART.md) - Quick start guide
- [API Documentation](backend/README.md#api-endpoints) - API reference

## License

MIT
