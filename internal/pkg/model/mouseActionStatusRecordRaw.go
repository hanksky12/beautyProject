package model

type MouseActionStatusRecordRaw struct {
	Base
	UserId        uint
	MouseActionId uint
	X             int64
	Y             int64
	Time          int64
}
