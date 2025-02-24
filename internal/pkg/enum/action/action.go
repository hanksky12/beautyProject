package action

import (
	"beautyProject/internal/pkg/enum"
)

type Action struct {
	*enum.Base
}

var (
	Start = &Action{&enum.Base{Number: 0, Name: "start"}}
	Stop  = &Action{&enum.Base{Number: 1, Name: "stop"}}
)
