package model

type HardwareStatusRecord struct {
	Base
	UserId     uint
	HardwareId uint
	Percent    float64
	Time       int64
	Processed  bool
}
