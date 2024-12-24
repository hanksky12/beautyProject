package request

import (
	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidations 注册所有自定义验证器
func RegisterCustomValidations(validate *validator.Validate) error {
	validators := map[string]validator.Func{

		"hardwareEnum": func(fl validator.FieldLevel) bool {
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
		},

		"exampleValidator": func(fl validator.FieldLevel) bool {
			value := fl.Field().String()
			// 示例验证逻辑：值必须为 "example" 或为空
			return value == "" || value == "example"
		},
	}

	// 逐个注册验证器
	for tag, fn := range validators {
		if err := validate.RegisterValidation(tag, fn); err != nil {
			return err
		}
	}

	return nil
}
