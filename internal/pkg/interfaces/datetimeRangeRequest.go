package interfaces

type IDatetimeRangeRequest interface {
	GetAllDateTime() (string, string, string, string)
}
