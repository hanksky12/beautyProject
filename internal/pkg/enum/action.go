package enum

type Action struct {
	*Base
}

var (
	Start = &Action{&Base{Number: 0, Name: "start"}}
	Stop  = &Action{&Base{Number: 1, Name: "stop"}}
)
