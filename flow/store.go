package flow

import (
	"context"

	"github.com/pkg/errors"
	"upper.io/db.v3/lib/sqlbuilder"

	"github.com/Sirupsen/logrus"
	"upper.io/db.v3/mysql"
)

// Configuration  configuration
type Configuration struct {
	Database string
	Host     string
	Port     int
	User     string
	Password string
}

// Storage storage
type Storage interface {
	GetFlow(ctx context.Context, id string) *Flow
}

// MyStorage storage
type MyStorage struct {
	configuration *Configuration
	settings      *mysql.ConnectionURL
	logger        *logrus.Logger
}

func defaultSettings(configuration *Configuration) *Configuration {
	if configuration.Port == 0 {
		configuration.Port = 3306
	}
	if configuration.Host == "" {
		configuration.Host = "localhost"
	}
	return configuration
}

func (p *MyStorage) conn() (sess sqlbuilder.Database, err error) {
	p.logger.Debugf("Connect to db with settings: %+v", p.settings)
	sess, err = mysql.Open(p.settings)
	if err != nil {
		err = errors.Wrap(err, "Failed to open db")
	}
	return
}

// NewStorage new storage
func NewStorage(logger *logrus.Logger, configuration *Configuration) (*MyStorage, error) {
	configuration = defaultSettings(configuration)
	storage := &MyStorage{
		configuration: configuration,
		settings: &mysql.ConnectionURL{
			Database: configuration.Database,
			Host:     configuration.Host,
			User:     configuration.User,
			Password: configuration.Password,
		},
		logger: logger,
	}
	return storage, nil
}

// GetFlow get one flow by id
func (p *MyStorage) GetFlow(id string) (*Flow, error) {
	sess, err := p.conn()
	if err != nil {
		return nil, err
	}

	defer sess.Close()
	var flow Flow
	err = sess.Collection("flow").Find("id", id).One(&flow)
	return &flow, err
}
