package model

type HardwareStatusRecordAverage struct {
	Base
	UserId     uint
	HardwareId uint
	Percent    float64
	Time       int64
}
