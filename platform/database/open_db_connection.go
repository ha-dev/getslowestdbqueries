package database

import (
	"github.com/ha-dev/getslowestdbqueries/queries"
	// "github.com/ha-dev/getslowestdbqueries/pkg/err"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.DbQueries // load queries from User model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		DbQueries: &queries.DbQueries{DB: db}, // from User model
	}, nil
}
