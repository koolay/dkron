package trigger

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

// NewTrigger instance trigger
func NewTrigger(logger *logrus.Logger, configuration *Configuration, kind string) (tr Trigger, err error) {
	switch strings.ToLower(kind) {
	case "rabbitmq":
		tr, err = newRabbitMQTrigger(logger, configuration)
	default:
		err = fmt.Errorf("Failed to NewTrigger. Not supported kind:%s", kind)
	}
	return
}
