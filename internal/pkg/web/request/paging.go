package request

type PagingReq struct {
	PerPage   int    `form:"per_page" binding:"required"`
	Page      int    `form:"page" binding:"required"`
	Sort      string `form:"sort" binding:"sortValidations"`
	SortOrder string `form:"sortOrder" binding:"required,oneof=asc desc"`
}
