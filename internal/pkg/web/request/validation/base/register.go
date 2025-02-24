package base

import (
	"beautyProject/internal/pkg/web/request/validation"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(validate *validator.Validate) error {
	validators := map[string]validator.Func{
		"hardwareValidations":      validation.HardwareValidations,
		"mouseActionValidations":   validation.MouseActionValidations,
		"sortValidations":          validation.SortValidations,
		"datetimeRangeValidations": validation.DatetimeRangeValidations,
		"rangeValidations":         validation.RangeValidations,
	}
	for tag, fn := range validators {
		if err := validate.RegisterValidation(tag, fn); err != nil {
			return err
		}
	}
	return nil
}
