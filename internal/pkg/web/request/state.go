package request

type HardwareReq struct {
	State    string `json:"state" binding:"required,oneof=start stop"`
	Hardware string `json:"hardware" binding:"required,oneof=cpu memory disk"`
}
