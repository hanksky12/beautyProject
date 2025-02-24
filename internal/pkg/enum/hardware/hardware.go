package hardware

import (
	"beautyProject/internal/pkg/enum"
)

type Hardware struct {
	*enum.Base
}

var (
	Cpu    = &Hardware{&enum.Base{Number: 1, Name: "cpu", ChineseName: "CPU"}}
	Disk   = &Hardware{&enum.Base{Number: 2, Name: "disk", ChineseName: "硬碟"}}
	Memory = &Hardware{&enum.Base{Number: 3, Name: "memory", ChineseName: "記憶體"}}
)

var Map = map[string]*Hardware{
	Cpu.Name:    Cpu,
	Disk.Name:   Disk,
	Memory.Name: Memory,
}
