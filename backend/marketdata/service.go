package marketdata

import (
	"encore.dev/storage/sqldb"
)

// Database instance for marketdata service
var db = sqldb.NewDatabase("marketdata", sqldb.DatabaseConfig{})

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
