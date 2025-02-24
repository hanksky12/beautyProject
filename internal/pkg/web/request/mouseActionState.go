package request

type MouseActionStateReq struct {
	Action string `form:"action" binding:"required,mouseActionValidations"`
	Token  string `form:"token" binding:"required"`
}
