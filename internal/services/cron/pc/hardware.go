package pc

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/repository"
)

type Hardware struct {
}

func (h *Hardware) Analyze(hardware *enum.Hardware, key string, value string, headers map[string]string, minutes int, repo repository.StatusRecord) {

}
