package marketdata

import (
	"context"

	"encore.dev/storage/sqldb"
)

// Service struct holds our dependencies
type Service struct {
	db *sqldb.Database
}

// Initialize the service
func initService() (*Service, error) {
	return &Service{
		db: sqldb.Named("marketdata"),
	}, nil
}

var svc = initService
