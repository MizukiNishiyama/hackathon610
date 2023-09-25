package usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

type RegisterUserUseCase struct {
	UserDao *dao.UserDao
}

func (uc *RegisterUserUseCase) Handle(user model.UserReqForHTTPPost) (model.User, error) {
	id := ulid.Make().String()
	userToInsert := model.User{
		Id:   id,
		Name: user.Name,
		Email:  user.Email,
	}

	err := uc.UserDao.Insert(userToInsert)
	if err != nil {
		return model.User{}, err
	}

	return userToInsert, nil
}
