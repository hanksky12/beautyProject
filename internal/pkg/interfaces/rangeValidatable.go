package interfaces

type IRangeValidatable interface {
	GetMaxMinValue(field string) (int64, int64)
}
