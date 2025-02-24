package model

type HardwareStatusRecordRaw struct {
	Base
	UserId     uint
	HardwareId uint
	Percent    float64
	Time       int64
}
