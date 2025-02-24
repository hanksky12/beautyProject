package request

type HardwareStateReq struct {
	State    string `json:"state" binding:"required,oneof=start stop"`
	Hardware string `json:"hardware" binding:"required,hardwareValidations"`
}
