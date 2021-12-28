package database

import (
	"fmt"

	"github.com/ha-dev/getslowestdbqueries/env"
	errors "github.com/ha-dev/getslowestdbqueries/pkg/err"

	//"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/postgres" // load pgx driver for PostgreSQL

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*gorm.DB, error) {
	// Define database connection settings.
	username := env.Db_user
	password := env.Db_pass
	dbName := env.Db_name
	dbHost := env.Db_host
	dbPort := env.Db_port
	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", dbHost, username, password, dbName, dbPort)

	// conn, err := gorm.Open("postgres", dbUri)
	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, errors.ErrorDatabaseConnection
	}

	return conn, nil

}
