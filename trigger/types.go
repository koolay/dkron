package trigger

// Configuration configuration
type Configuration struct {
	RabbitMQ RabbitMQConfiguration
}

// RabbitMQConfiguration rabbitmq
type RabbitMQConfiguration struct {
	URI          string
	ExchangeName string
	QueueName    string
	BindKey      string
	ConsumerTag  string
	Topics       []string
}

// JobMessage received job from trigger
type JobMessage struct {
	ID             string `json:"id"`
	ProjectCode    string `json:"project_code"`
	FlowID         string `json:"flow_id"`
	FlowInstanceID string `json:"flow_instance_id"`
}
