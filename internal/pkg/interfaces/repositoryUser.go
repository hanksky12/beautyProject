package interfaces

import "beautyProject/internal/pkg/model"

type IRepoUser interface {
	FindByName(name string) (*model.User, bool)
	Add(user *model.User) error
	Remove(id int)
}
