package kafkaUtil

type Topic struct {
	Name        string `mapstructure:"name"`
	Partition   int    `mapstructure:"partition"`
	Replication int    `mapstructure:"replication"`
	Type        string `mapstructure:"type"`
}

type Broker struct {
	Address string `mapstructure:"address"`
}

type KafkaConfig struct {
	Brokers []Broker `mapstructure:"brokers"`
	Topics  []Topic  `mapstructure:"topics"`
}
