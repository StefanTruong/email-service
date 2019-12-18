package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Options for the postgres database connection
type Options struct {
	// Host name of the database
	Host string

	// Port of the postgres database
	Port int

	// User to use for the connection
	User string
	// Password of the user
	// Password string

	// Database to connect to
	Database string
}

func Connect(opts Options) (*sql.DB, error) {
	// Creating a ConnectionString
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		opts.Host, opts.Port, opts.User, opts.Database)
	log.Println(psqlInfo)

	// open a connection to database
	return sql.Open("postgres", psqlInfo)
}
