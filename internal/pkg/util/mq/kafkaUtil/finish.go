package kafkaUtil

func CloseWriter() {
	for _, writer := range WriterMap {
		writer.Close()
	}
}
