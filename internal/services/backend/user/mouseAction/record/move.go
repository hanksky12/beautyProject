package record

import (
	"beautyProject/internal/services/backend/common/record"
)

type Move struct {
	Recorder record.Recorder
}

func (m *Move) Work(userId string, msg []byte) {
	m.Recorder.Work(userId, msg)
}
