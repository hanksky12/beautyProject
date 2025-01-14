package validation

import "github.com/go-playground/validator/v10"

func HardwareValidations(fl validator.FieldLevel) bool {
	hardware := fl.Field().String()
	if hardware == "" {
		return true
	}
	switch hardware {
	case "cpu", "memory", "disk":
		return true
	default:
		return false
	}
}
