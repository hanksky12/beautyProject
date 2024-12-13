package initApp

type Options struct {
	KafkaOptions *KafkaOptions
	LogOptions   *LogOptions
}

type KafkaOptions struct {
	IsKafka        bool
	IsCreateWriter bool
}

type LogOptions struct {
	FilePath string
}
