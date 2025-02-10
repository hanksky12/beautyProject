package user

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/interfaces"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/str"
	"beautyProject/internal/pkg/util/web"
	"beautyProject/internal/pkg/web/request"
	log "github.com/sirupsen/logrus"
)

type User struct {
}

func (u *User) Register(jwt interfaces.ISetCookie, userRepo interfaces.IRepoUser, req request.UserReq) dto.Msg {
	//panic("no value for $USER")
	user, isFind := userRepo.FindByName(req.UserName)
	if isFind == true {
		msg := dto.Msg{Success: false, Message: "用戶名被註冊"}
		log.Infof("%v", msg)
		return msg
	}
	user = &model.User{Name: req.UserName, Password: str.HashPassword(req.UserPassword), AuthorizationLevel: "user"}
	err := userRepo.Add(user)
	if err != nil {
		msg := dto.Msg{Success: false, Message: "註冊失敗"}
		log.Infof("%v", msg)
		return msg
	}
	err = jwt.SetCookie(user.ID)
	if err != nil {
		msg := dto.Msg{Success: false, Message: "設定cookie失敗"}
		log.Infof("%v", msg)
		return msg
	}
	msg := dto.Msg{Success: true, Message: "註冊成功"}
	log.Infof("%v", msg)
	return msg

}

func (u *User) Login(jwt interfaces.ISetCookie, userRepo interfaces.IRepoUser, req request.UserReq) dto.Msg {
	user, isFind := userRepo.FindByName(req.UserName)
	if isFind == false {
		msg := dto.Msg{Success: false, Message: "用戶未註冊"}
		log.Infof("%v", msg)
		return msg
	}
	if user.Password != str.HashPassword(req.UserPassword) {
		msg := dto.Msg{Success: false, Message: "密碼錯誤"}
		log.Infof("%v", msg)
		return msg
	}
	err := jwt.SetCookie(user.ID)
	if err != nil {
		log.Info(err)
		msg := dto.Msg{Success: false, Message: "設定cookie失敗"}
		log.Infof("%v", msg)
		return msg
	}
	msg := dto.Msg{Success: true, Message: "登入成功"}
	log.Infof("%v", msg)
	return msg
}

func (u *User) Logout(jwt *web.Jwt) dto.Msg {
	jwt.UnsetCookie()
	msg := dto.Msg{Success: true, Message: "登出成功"}
	log.Infof("%v", msg)
	return msg
}
