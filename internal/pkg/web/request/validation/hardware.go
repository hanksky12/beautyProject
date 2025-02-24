package validation

import (
	enumHardware "beautyProject/internal/pkg/enum/hardware"
	"github.com/go-playground/validator/v10"
)

func HardwareValidations(fl validator.FieldLevel) bool {
	hardware := fl.Field().String()
	if hardware == "" {
		return true
	}
	for _, value := range enumHardware.Map {
		if hardware == value.ChineseName {
			return true
		}
	}
	return false

	//switch hardware {
	//case enumHardware.Cpu.ChineseName, enumHardware.Disk.ChineseName, enumHardware.Memory.ChineseName:
	//	return true
	//default:
	//	return false
	//}
}
