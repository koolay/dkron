// Package db provides database access
package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

// SqliteDatabase sqlite implement
type SqliteDatabase struct {
	BaseDatabase
}

// NewSqliteDatabase init sqlite Database instance
func NewSqliteDatabase(logger *logrus.Logger, configuration Configuration) (database Database, err error) {
	if configuration.ConnectionURL == "" {
		configuration.ConnectionURL = ""
	}
	database = &SqliteDatabase{
		BaseDatabase{logger: logger, configuration: configuration},
	}

	settings, urlErr := sqlite.ParseURL(configuration.ConnectionURL)
	if err != nil {
		return database, urlErr
	}

	conn, err := sqlite.Open(settings)
	if err == nil {
		defer conn.Close()
		errPing := conn.Ping()
		if errPing != nil {
			err = errors.Wrapf(errPing, "Failed to ping: %s", configuration.ConnectionURL)
		}
	}

	return
}

// Open new connection
func (d *SqliteDatabase) Open() (conn sqlbuilder.Database, err error) {
	return sqlite.Open(d.connectionURL)
}
