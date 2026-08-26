# Ent Migration Guide

## Overview

This project uses **Ent** for ORM and migrations. Ent provides automatic schema generation and migration support.

## Schema Files

Schema files define your database entities:

- `backend/marketdata/ent/schema/` - Market Data schemas
- `backend/news/ent/schema/` - News schemas

## Workflow

### 1. Define Schema

Edit schema files in `ent/schema/`:

```go
// backend/marketdata/ent/schema/stock.go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type Stock struct {
    ent.Schema
}

func (Stock) Fields() []ent.Field {
    return []ent.Field{
        field.String("stock_code").Unique(),
        field.String("stock_name"),
    }
}
```

### 2. Generate Code

Generate Ent code from schemas:

```bash
cd backend
make generate
```

Or manually:

```bash
cd marketdata
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

cd ../news
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
```

This creates:
- Entity structs
- Query builders
- Mutation builders
- Client code

### 3. Run Migrations

Encore handles migrations automatically when you run:

```bash
make dev
```

Or manually:

```bash
encore db migrate --schema marketdata
encore db migrate --schema news
```

## Common Schema Patterns

### Basic Field Types

```go
field.String("name")
field.Int("age")
field.Bool("active")
field.Time("created_at")
field.Float("price")
field.JSON("metadata", map[string]interface{}{})
```

### Field Modifiers

```go
field.String("email").
    Unique().              // Unique constraint
    NotEmpty().           // Not null + not empty
    MaxLen(255).          // Max length
    Optional().           // Nullable
    Default("default").   // Default value
    Immutable().          // Cannot be updated
    Comment("User email") // SQL comment
```

### Relationships

```go
// One-to-Many
edge.To("summaries", StockSummary.Type)

// Many-to-One
edge.From("stock", Stock.Type).
    Ref("summaries").
    Required().
    Unique()
```

### Indexes

```go
func (Stock) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("stock_code").Unique(),
        index.Fields("created_at", "status"),
    }
}
```

### Enums

```go
field.Enum("status").
    Values("pending", "processing", "completed", "failed").
    Default("pending")
```

## Using Ent Client

### Initialize Client

```go
import "encore.app/marketdata/ent"

client := ent.NewClient(ent.Driver(drv))
defer client.Close()
```

### Create

```go
stock, err := client.Stock.
    Create().
    SetStockCode("BBCA").
    SetStockName("Bank Central Asia").
    Save(ctx)
```

### Query

```go
// Get by ID
stock, err := client.Stock.Get(ctx, id)

// Query with filters
stocks, err := client.Stock.
    Query().
    Where(stock.StockCodeHasPrefix("BB")).
    Limit(10).
    All(ctx)

// Count
count, err := client.Stock.Query().Count(ctx)
```

### Update

```go
err := client.Stock.
    UpdateOneID(id).
    SetStockName("New Name").
    Save(ctx)
```

### Delete

```go
err := client.Stock.DeleteOneID(id).Exec(ctx)
```

### Transactions

```go
tx, err := client.Tx(ctx)
if err != nil {
    return err
}

defer func() {
    if v := recover(); v != nil {
        tx.Rollback()
        panic(v)
    }
}()

// ... perform operations ...

if err := tx.Commit(); err != nil {
    return err
}
```

## Migration Commands

### View Migration Status

```bash
encore db migrations list --schema marketdata
```

### Create New Migration

After modifying schemas:

```bash
make generate
encore run  # Auto-creates and applies migrations
```

### Rollback Migration

```bash
encore db migrate --schema marketdata --down 1
```

### Reset Database

```bash
make db-reset
```

## Best Practices

1. **Always generate after schema changes**
   ```bash
   make generate
   ```

2. **Use transactions for multiple operations**

3. **Add indexes for frequently queried fields**

4. **Use comments for documentation**
   ```go
   field.String("email").Comment("User email address")
   ```

5. **Validate data in schema**
   ```go
   field.String("email").
       NotEmpty().
       MaxLen(255).
       Match(regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"))
   ```

6. **Use hooks for business logic**
   ```go
   func (Stock) Hooks() []ent.Hook {
       return []ent.Hook{
           hook.On(
               func(next ent.Mutator) ent.Mutator {
                   // Hook logic
                   return next
               },
               ent.OpCreate|ent.OpUpdate,
           ),
       }
   }
   ```

## Troubleshooting

### Schema Generation Fails

```bash
# Clean and regenerate
rm -rf marketdata/ent/
rm -rf news/ent/
make generate
```

### Migration Conflicts

```bash
# Reset database
make db-reset
make dev
```

### Type Errors After Schema Change

```bash
# Regenerate and rebuild
make generate
go build ./...
```

## Resources

- [Ent Documentation](https://entgo.io/docs/getting-started)
- [Schema Definition](https://entgo.io/docs/schema-def)
- [CRUD Operations](https://entgo.io/docs/crud)
- [Migrations](https://entgo.io/docs/migrate)
