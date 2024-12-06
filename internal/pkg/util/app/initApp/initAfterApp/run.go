package initAfterApp

import (
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
)

func Run(createWriter bool) {
	if createWriter {
		kafkaUtil.CloseWriter()
	}
}
