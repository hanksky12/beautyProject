package validation

import "github.com/go-playground/validator/v10"

func SortValidations(fl validator.FieldLevel) bool {
	sort := fl.Field().String()
	if sort == "" {
		return true
	}
	switch sort {
	case "hardware_name", "time", "percent", "mouse_action_name", "x", "y":
		return true
	default:
		return false
	}
}
