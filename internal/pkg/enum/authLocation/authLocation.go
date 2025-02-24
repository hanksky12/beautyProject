package authLocation

import (
	"beautyProject/internal/pkg/enum"
)

type AuthLocation struct {
	*enum.Base
}

var (
	Cookie      = &AuthLocation{&enum.Base{Number: 0, Name: "cookie"}}
	QueryParams = &AuthLocation{&enum.Base{Number: 1, Name: "query_params"}}
)
