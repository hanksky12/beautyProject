package entity

type HardwareStatusRecordWithUser struct {
	HardwareStatusRecordID uint
	UserId                 uint
	HardwareId             uint
	Percent                float64
	Time                   int64
	Processed              bool
}
