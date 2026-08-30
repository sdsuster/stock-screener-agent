package news

import (
	"encore.dev/storage/sqldb"
)

// Database instance for news service
var db = sqldb.NewDatabase("news", sqldb.DatabaseConfig{})

// Service struct holds our dependencies
type Service struct {
	db *sqldb.Database
}

// Initialize the service
func initService() (*Service, error) {
	return &Service{
		db: db,
	}, nil
}

var svc = initService
