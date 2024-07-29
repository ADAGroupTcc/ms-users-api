package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Environments define the environment variables
type Environments struct {
	ApiPort string `envconfig:"PORT" default:"8081"`

	DBUri  string `envconfig:"MONGODB_URI"`
	DBName string `envconfig:"MONGODB_DBNAME"`

	KafkaBrokers     string `envconfig:"KAFKA_BROKERS" default:"localhost"`
	KafkaTopicOutput string `envconfig:"KAFKA_TOPIC_OUTPUT" default:"output"`
	KafkaConsumerId  string `envconfig:"KAFKA_CONSUMER_ID" default:"0"`
}

// LoadEnvVars load the environment variables
func LoadEnvVars() (*Environments, error) {
	godotenv.Load()
	c := &Environments{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}
