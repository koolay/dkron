package flow

import (
	"context"

	"github.com/Sirupsen/logrus"
	"github.com/victorcoder/dkron/db"
)

// Storage storage
type Storage interface {
	GetFlow(ctx context.Context, id string) *Flow
}

// MyStorage storage
type MyStorage struct {
	database db.Database
	logger   *logrus.Logger
}

// NewStorage new storage
func NewStorage(logger *logrus.Logger, database db.Database) (*MyStorage, error) {
	storage := &MyStorage{
		database: database,
		logger:   logger,
	}
	return storage, nil
}

// GetFlow get one flow by id
func (p *MyStorage) GetFlow(id string) (*Flow, error) {
	sess, err := p.database.Open()
	if err != nil {
		return nil, err
	}
	defer sess.Close()
	var flow Flow
	err = sess.Collection("flow").Find("id", id).One(&flow)
	return &flow, err
}
