package base

var personalStatusRecord map[string]hardware

type hardware struct {
	Name map[string]state
}

type state struct {
	IsWorking bool
	StopChan  chan bool
}

/*
personalStatusRecord
├── user1
│   ├── cpu: { IsWorking: true, StopChan: channel }
│   └── disk: { IsWorking: false, StopChan: channel }
└── user2
├── cpu: { IsWorking: true, StopChan: channel }
├── disk: { IsWorking: false, StopChan: channel }
└── memory: { IsWorking: true, StopChan: channel }
*/
