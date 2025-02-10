package entity

type StatusRecordWithUser struct {
	StatusRecordID uint
	UserId         uint
	HardwareId     uint
	Percent        float64
	Time           int64
	Processed      bool
}
