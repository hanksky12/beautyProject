package model

type User struct {
	Base
	Name               string
	Password           string
	AuthorizationLevel string
}
