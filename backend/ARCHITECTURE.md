# Architecture Overview

## System Design

```
┌─────────────────────────────────────────────────────────────┐
│                         Client Layer                         │
│  (Web/Mobile Apps, CLI, External Services)                  │
└────────────────────┬────────────────────────────────────────┘
                     │
                     │ HTTP/REST
                     │
┌────────────────────▼────────────────────────────────────────┐
│                      Gateway Service                         │
│  • API Routing                                               │
│  • Request Aggregation                                       │
│  • Health Checks                                             │
└───────────────┬─────────────────────┬───────────────────────┘
                │                     │
        ┌───────▼──────┐      ┌──────▼────────┐
        │  Market Data │      │  News Service │
        │   Service    │      │               │
        └───────┬──────┘      └───────┬───────┘
                │                     │
        ┌───────▼──────┐      ┌──────▼────────┐
        │   PostgreSQL │      │   PostgreSQL  │
        │   Database   │      │   + pgvector  │
        │              │      │   Database    │
        └──────────────┘      └───────────────┘
```

## Service Details

### Gateway Service
**Purpose**: Unified API entry point and request orchestration

**Responsibilities**:
- Route requests to appropriate services
- Aggregate data from multiple services
- Provide health checks
- Cross-service search

**Endpoints**:
- `GET /health` - System health check
- `GET /api/stocks/:code` - Get stock with related news
- `POST /api/search` - Search across services

**Dependencies**: Market Data, News

---

### Market Data Service
**Purpose**: Stock ticker and market data management

**Responsibilities**:
- Manage stock tickers (CRUD)
- Store daily stock summaries
- Historical data queries
- Market data aggregation

**Endpoints**:
- `GET /marketdata/stocks` - List all stocks
- `GET /marketdata/stocks/:code` - Get single stock
- `POST /marketdata/stocks` - Create/update stock
- `POST /marketdata/summaries` - Query summaries

**Database Schema**:
```sql
-- stocks table
CREATE TABLE stocks (
    stock_code VARCHAR(20) PRIMARY KEY,
    stock_name VARCHAR(255) NOT NULL
);

-- stock_summaries table
CREATE TABLE stock_summaries (
    id SERIAL PRIMARY KEY,
    stock_code VARCHAR(20) REFERENCES stocks(stock_code),
    date DATE NOT NULL,
    previous BIGINT,
    open_price BIGINT,
    high BIGINT,
    low BIGINT,
    close BIGINT,
    volume BIGINT,
    value BIGINT,
    -- ... more fields
    UNIQUE(stock_code, date)
);
```

---

### News Service
**Purpose**: News article storage and semantic search

**Responsibilities**:
- Store news articles
- Chunk content for embeddings
- Generate vector embeddings
- Semantic similarity search
- Manage embedding processing status

**Endpoints**:
- `POST /news` - Create news article
- `GET /news/:id` - Get single article
- `POST /news/list` - List articles
- `POST /news/search` - Semantic search
- `GET /news/:id/chunks` - Get article chunks

**Database Schema**:
```sql
-- news table
CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    url VARCHAR(2048) UNIQUE NOT NULL,
    published_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- news_chunks table
CREATE TABLE news_chunks (
    id SERIAL PRIMARY KEY,
    news_id INTEGER REFERENCES news(id),
    embedding_type VARCHAR(50),
    chunk_index INTEGER,
    content TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    embedding vector(1024),  -- pgvector
    UNIQUE(news_id, chunk_index)
);

-- Vector similarity index
CREATE INDEX ON news_chunks USING ivfflat (embedding vector_cosine_ops);
```

---

## Technology Stack

### Framework
- **Encore**: Modern Go framework for microservices
  - Built-in service discovery
  - Automatic API documentation
  - Database management
  - Distributed tracing
  - Local development dashboard

### ORM
- **Ent**: Entity framework for Go
  - Code generation
  - Type-safe queries
  - Automatic migrations
  - Relationship management
  - Schema-first design

### Database
- **PostgreSQL**: Primary database
  - ACID compliance
  - JSON support
  - Full-text search
  
- **pgvector**: Vector similarity extension
  - 1024-dimensional embeddings
  - Cosine similarity search
  - IVFFlat indexing

---

## Data Flow

### Stock Data Ingestion
```
External API → Python Crawler → Market Data Service → PostgreSQL
                                      ↓
                                  Gateway API
                                      ↓
                                   Clients
```

### News Processing
```
News Source → Python Crawler → News Service → PostgreSQL
                                    ↓
                              Chunking Worker
                                    ↓
                            Embedding Worker (Python)
                                    ↓
                              Update Embeddings
                                    ↓
                          Vector Search Available
```

### Search Query
```
Client → Gateway → News Service → pgvector similarity search
                       ↓
                  Ranked Results
```

---

## Communication Patterns

### Synchronous (HTTP/REST)
- Client ↔ Gateway
- Gateway ↔ Services
- Services ↔ Database

**Use Cases**: 
- CRUD operations
- Queries
- Real-time data retrieval

### Asynchronous (Future)
- Message Queue (RabbitMQ/Kafka)
- Background jobs
- Event-driven updates

**Use Cases**:
- Embedding generation
- Data synchronization
- Bulk processing

---

## Scalability Considerations

### Horizontal Scaling
- Each service can scale independently
- Encore handles service discovery
- Load balancing via container orchestration

### Database Scaling
- Read replicas for queries
- Partitioning by date (stock_summaries)
- Connection pooling

### Caching Strategy
- Redis for frequently accessed data
- Cache at gateway level
- TTL-based invalidation

---

## Security

### Authentication (Future)
- JWT tokens
- API keys for external services
- Service-to-service auth

### Authorization
- Role-based access control
- Service-level permissions

### Data Protection
- HTTPS/TLS encryption
- Database encryption at rest
- Sensitive data masking

---

## Monitoring & Observability

### Built-in (Encore)
- Request tracing
- Performance metrics
- Error tracking
- Service health

### Custom Metrics
- Stock data freshness
- Embedding processing rate
- Search query performance
- API latency

### Logging
- Structured logging (JSON)
- Log levels (debug, info, warn, error)
- Centralized log aggregation

---

## Deployment

### Local Development
```bash
make dev  # Run all services
```

### Docker
```bash
docker-compose up  # Database
encore build docker  # Build image
```

### Encore Cloud
```bash
encore deploy  # Deploy to cloud
```

### Kubernetes (Future)
- Helm charts
- Auto-scaling
- Rolling updates
- Health checks

---

## Future Enhancements

### Phase 1 (Current)
- ✅ Basic microservices
- ✅ Ent ORM integration
- ✅ Vector search support

### Phase 2
- [ ] Authentication & authorization
- [ ] Redis caching
- [ ] Rate limiting
- [ ] API versioning

### Phase 3
- [ ] Message queue integration
- [ ] Real-time updates (WebSocket)
- [ ] Advanced analytics
- [ ] Machine learning integration

### Phase 4
- [ ] Multi-region deployment
- [ ] Event sourcing
- [ ] GraphQL API
- [ ] Mobile SDKs

---

## Performance Targets

| Metric | Target | Notes |
|--------|--------|-------|
| API Latency (p95) | < 100ms | Gateway endpoints |
| Database Query | < 50ms | Simple queries |
| Vector Search | < 200ms | Top 10 results |
| Throughput | 1000 rps | Per service |
| Availability | 99.9% | Monthly uptime |

---

## Development Workflow

1. **Schema Design** - Define Ent schemas
2. **Generate Code** - Run `make generate`
3. **Implement Logic** - Add business logic
4. **Write Tests** - Unit and integration tests
5. **Run Locally** - `make dev`
6. **Deploy** - `encore deploy`

---

## Resources

- [Encore Documentation](https://encore.dev/docs)
- [Ent Documentation](https://entgo.io/docs)
- [pgvector Documentation](https://github.com/pgvector/pgvector)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
