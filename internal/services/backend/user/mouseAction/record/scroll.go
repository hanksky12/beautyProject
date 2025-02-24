package record

import (
	"beautyProject/internal/services/backend/common/record"
)

type Scroll struct {
	Recorder record.Recorder
}

func (s *Scroll) Work(userId string, msg []byte) {
	s.Recorder.Work(userId, msg)
}
