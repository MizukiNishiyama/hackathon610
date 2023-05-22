package usecase

import (
	"sourse/dao"
	"sourse/model"
)

type SearchUserUseCase struct {
	UserDao *dao.UserDao
}

func (uc *SearchUserUseCase) Handle(name string) ([]model.User, error) {
	return uc.UserDao.FindByName(name)
}
