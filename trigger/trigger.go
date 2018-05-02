package trigger

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

var (
	jobMessagePool = sync.Pool{New: func() interface{} {
		return &JobMessage{}
	}}
)

type AbstractTrigger struct {
	ID     string
	Logger *logrus.Logger
	Kind   string
}

type Trigger interface {
	Initialize() error
	Start() error
	Stop() error
}
