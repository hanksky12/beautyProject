package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
)

type User struct {
}

func (u *User) FindByName(name string) (*model.User, bool) {
	user := &model.User{}
	result := sql.Db.Where("name = ?", name).Find(user)
	if result.RowsAffected == 0 {
		//log.Info("No user found with name:", name)
		return nil, false
	}
	//log.Info("User found")
	return user, true
}

func (u *User) Add(user *model.User) error {
	return sql.Db.Create(user).Error
}

func (u *User) Remove(id int) {
	sql.Db.Delete(&model.User{}, id)
}
