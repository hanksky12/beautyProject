package enum

type Hardware struct {
	*Base
}

var (
	Cpu    = &Hardware{&Base{Number: 0, Name: "cpu"}}
	Disk   = &Hardware{&Base{Number: 1, Name: "disk"}}
	Memory = &Hardware{&Base{Number: 2, Name: "memory"}}
)
