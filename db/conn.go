// Package db provides database access
package db

import (
	"context"

	"github.com/Sirupsen/logrus"
	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Configuration  configuration
type Configuration struct {
	ConnectionURL string
}

// Database provider db abstract
type Database interface {
	Open() (sqlbuilder.Database, error)
}

// BaseDatabase base struct
type BaseDatabase struct {
	Driver        string
	ctx           context.Context
	logger        *logrus.Logger
	configuration Configuration
	connectionURL db.ConnectionURL
}
