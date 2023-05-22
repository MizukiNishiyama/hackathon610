package usecase

import (
	"github.com/oklog/ulid/v2"
	"sourse/dao"
	"sourse/model"
)

type RegisterUserUseCase struct {
	UserDao *dao.UserDao
}

func (uc *RegisterUserUseCase) Handle(user model.UserReqForHTTPPost) (model.User, error) {
	id := ulid.Make().String()
	userToInsert := model.User{
		Id:   id,
		Name: user.Name,
		Age:  user.Age,
	}

	err := uc.UserDao.Insert(userToInsert)
	if err != nil {
		return model.User{}, err
	}

	return userToInsert, nil
}
