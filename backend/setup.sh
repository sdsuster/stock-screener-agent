#!/usr/bin/env bash

set -e

echo "🚀 Setting up Stock Screener Microservices..."
echo ""

# Check if Encore is installed
if ! command -v encore &> /dev/null; then
    echo "📦 Installing Encore CLI..."
    npm install -g encore
else
    echo "✅ Encore CLI already installed"
fi

# Check if Docker is running
if ! docker info &> /dev/null; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
else
    echo "✅ Docker is running"
fi

# Navigate to backend directory
cd "$(dirname "$0")"

# Install Go dependencies
echo ""
echo "📦 Installing Go dependencies..."
go mod download

# Start database
echo ""
echo "🗄️  Starting PostgreSQL with pgvector..."
docker-compose up -d

# Wait for database
echo "⏳ Waiting for database to be ready..."
sleep 5

# Generate Ent code
echo ""
echo "🔨 Generating Ent code..."
cd marketdata
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
cd ../news
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
cd ..

echo ""
echo "✅ Setup complete!"
echo ""
echo "🎯 Next steps:"
echo "   1. Run services: make dev"
echo "   2. Open dashboard: http://localhost:9400"
echo "   3. Test health: http://localhost:4000/health"
echo ""
echo "📚 Documentation:"
echo "   - README.md - Full documentation"
echo "   - QUICKSTART.md - Quick start guide"
echo "   - MIGRATIONS.md - Database migrations"
echo ""
echo "🛠️  Useful commands:"
echo "   make dev          - Run services"
echo "   make test         - Run tests"
echo "   make db-down      - Stop database"
echo "   make help         - Show all commands"
echo ""
