// Package db provides database access
package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

// MysqlDatabase mysql implement
type MysqlDatabase struct {
	BaseDatabase
}

// NewMysqlDatabase init mysql Database instance
func NewMysqlDatabase(logger *logrus.Logger, configuration Configuration) (database Database, err error) {
	if configuration.ConnectionURL == "" {
		configuration.ConnectionURL = ""
	}
	database = &MysqlDatabase{
		BaseDatabase{logger: logger, configuration: configuration},
	}

	settings, urlErr := mysql.ParseURL(configuration.ConnectionURL)
	if err != nil {
		return database, urlErr
	}

	conn, err := mysql.Open(settings)
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
func (d *MysqlDatabase) Open() (conn sqlbuilder.Database, err error) {
	return mysql.Open(d.connectionURL)
}
