package news

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
		db: sqldb.Named("news"),
	}, nil
}

var svc = initService
