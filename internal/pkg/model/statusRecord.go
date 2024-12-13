package model

type StatusRecord struct {
	Base
	UserId     uint
	HardwareId uint
	Percent    float64
	Time       int64
	Processed  bool
}

type MiniStatusRecord struct {
	ID         uint
	UserId     uint
	HardwareId uint
	Percent    float64
	Time       int64
}
