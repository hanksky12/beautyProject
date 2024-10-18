package user

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/pkg/web/request"
	log "github.com/sirupsen/logrus"
)

type User struct {
}

func (u *User) Register(req request.UserReq) dto.Msg {
	//panic("no value for $USER")
	userRepo := repository.User{}
	//userRepo.Add(req)
	//userRepo.Delete(1)
	user := userRepo.Find(req.UserName)
	if user.RowsAffected != 0 {

	}
	msg := dto.Msg{Success: true, Message: "註冊成功"}
	log.Infof("%v", msg)
	return msg

}

func (u *User) Login(req request.UserReq) {

}

func (u *User) Logout() {

}
