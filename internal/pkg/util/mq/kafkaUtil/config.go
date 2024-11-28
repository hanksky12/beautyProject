package kafkaUtil

type Topic struct {
	Name        string `mapstructure:"name"`
	Partition   int    `mapstructure:"partition"`
	Replication int    `mapstructure:"replication"`
}

type Broker struct {
	Address string `mapstructure:"address"`
}

type KafkaConfig struct {
	Brokers []Broker `mapstructure:"brokers"`
	Topics  []Topic  `mapstructure:"topics"`
}
