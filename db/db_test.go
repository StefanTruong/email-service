package db

import (
	"github.com/StefanTruong/email-service/env"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Run("should return an error if the connection options don't match a database", func(t *testing.T) {
		opts, err := env.Parse()
		if err != nil {
			t.Error(err)
		}

		opts.Port = 4444
		_, err = Connect(opts)

		if err == nil {
			t.Fatalf("error is %#v", err)
		}
	})

	t.Run("should return the database connection if the options are correct", func (t *testing.T) {
		opts, err := env.Parse()
		if err != nil {
			t.Error(err)
		}
		conn, err := Connect(opts)

		if err != nil {
			t.Fatalf("error connecting to db: %s", err.Error())
		}

		if conn == nil {
			t.Fatalf("database connection is nil")
		}
	})
}