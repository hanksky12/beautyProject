package model

type StatusRecordAverage struct {
	Base
	UserId     uint
	HardwareId uint
	Percent    float64
	Time       int64
}
