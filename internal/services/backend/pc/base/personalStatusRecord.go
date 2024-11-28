package base

var personalStatusRecord map[string]hardware

type hardware struct {
	Name map[string]state
}

type state struct {
	IsWorking bool
	StopChan  chan bool
}
