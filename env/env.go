package env

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

// Parse returns a struct containing all important environment variables
func Parse() (Env, error) {
	var e Env = Env{}
	if err := env.Parse(&e); err != nil {
		return e, fmt.Errorf("parse environment variables: %w", err)
	}
	return e, nil
}

// Env contains the values of important environment variables
type Env struct {
	// Hostname is the connection string used to connect to the Postgres DB.
	// Example value: "localhost"
	Hostname string `env:"PG_HOST,required"`

	// Port that the Postgres DB listens to.
	Port int `env:"PG_PORT,required"`

	// User is the Username for the DB
	User string `env:"PG_USER,required"`

	// Password
	Password string `env:"PG_PASSWORD,required"`

	// Database is the Username for the DB
	Database string `env:"PG_DATABASE,required"`
}