package enum

type Hardware struct {
	*Base
}

var (
	Cpu    = &Hardware{&Base{Number: 1, Name: "cpu"}}
	Disk   = &Hardware{&Base{Number: 2, Name: "disk"}}
	Memory = &Hardware{&Base{Number: 3, Name: "memory"}}
)
