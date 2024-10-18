package user

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/web/request"
	log "github.com/cihub/seelog"
)

type User struct {
	ReqId string
}

func (u *User) Register(req request.UserReq) dto.Msg {
	msg := dto.Msg{Success: true, Message: "註冊成功"}
	log.Infof("%v %v", u.ReqId, msg)
	//panic("no value for $USER")
	return msg

}

func (u *User) Login(req request.UserReq) {

}

func (u *User) Logout() {

}
