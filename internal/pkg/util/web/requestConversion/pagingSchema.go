package requestConversion

type PagingSchema struct {
	PerPage   int
	Page      int
	Sort      string
	SortOrder string
}
