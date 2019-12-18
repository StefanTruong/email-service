package db

import (
	"database/sql"
	"fmt"
	"github.com/StefanTruong/email-service/env"
	_ "github.com/lib/pq"
)

func Connect(opts env.Env) (*sql.DB, error) {
	// Creating a ConnectionString
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		opts.Hostname, opts.Port, opts.User, opts.Database)

	// open a connection to database
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec("SELECT current_date")
	if err != nil {
		return nil, fmt.Errorf("failed test exec: %w", err)
	}

	return db, nil
}
