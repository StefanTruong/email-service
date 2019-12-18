package db

import (
	"testing"
)

func TestConnect(t *testing.T) {
	t.Run("should return an error if the connection options don't match a database", func(t *testing.T) {
		_, err := Connect(Options{
			Host:     "localhost",
			Port:     5894,
			User:     "postgres",
			Database: "postgres",
		})

		if err == nil {
			t.Fatalf("error is %#v", err)
		}
	})

	t.Run("should return the database connection if the options are correct", func (t *testing.T) {
		conn, err := Connect(Options{
			Host:     "/var/run/postgresql",
			Port:     5432,
			User:     "stefan",
			Database: "email_service",
		})

		if err != nil {
			t.Fatalf("error connecting to db: %s", err.Error())
		}

		if conn == nil {
			t.Fatalf("database connection is nil")
		}
	})
}