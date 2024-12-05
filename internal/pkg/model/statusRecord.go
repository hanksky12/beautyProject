package model

type StatusRecord struct {
	Base
	UserId     uint
	HardwareId uint
	Percent    float64
	Time       int64
	Min        int
}
