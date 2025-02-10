package interfaces

import "github.com/go-co-op/gocron"

type IEnvironment interface {
	InProduction(scheduler *gocron.Scheduler)
	InOther(scheduler *gocron.Scheduler)
}
