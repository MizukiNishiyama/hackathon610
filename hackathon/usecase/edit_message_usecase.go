package usecase

import (
	"hackathon/dao"
)

type EditMessageUsecase struct {
	MessageDao *dao.MessageDao
}

func (uc *EditMessageUsecase) EditMessage(id, content string) error {
	return uc.MessageDao.EditMessage(id, content)
}
