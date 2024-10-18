package repository

import (
	"beautyProject/internal/pkg/db/sql"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/web/request"
)

type User struct{}

func (u *User) Find(name string) model.User {
	user := model.User{Name: name}
	result := sql.Db.First(&user)
	return result
}

func (u *User) Add(req request.UserReq) {
	user := model.User{Name: req.UserName, Password: req.UserPassword, AuthorizationLevel: "user"}
	sql.Db.Create(&user)
}

func (u *User) Delete(id int) {
	sql.Db.Delete(&model.User{}, id)
}
