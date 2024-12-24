package request

type PagingReq struct {
	PerPage   int    `json:"per_page" binding:"required"`
	Page      int    `json:"page" binding:"required"`
	Sort      string `json:"sort" binding:"required,oneof=cpu memory disk"`
	SortOrder string `json:"sort_order" binding:"required,oneof=asc desc"`
}
