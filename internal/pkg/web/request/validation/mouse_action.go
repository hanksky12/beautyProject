package validation

import (
	"beautyProject/internal/pkg/enum/mouseAction"
	"github.com/go-playground/validator/v10"
)

func MouseActionValidations(fl validator.FieldLevel) bool {
	action := fl.Field().String()
	if action == "" {
		return true
	}
	for _, value := range mouseAction.Map {
		if action == value.ChineseName {
			return true
		}
	}
	return false

	//switch action {
	//case mouseAction.Click.ChineseName, mouseAction.Move.ChineseName, mouseAction.Scroll.ChineseName:
	//	return true
	//default:
	//	return false
	//}
}
