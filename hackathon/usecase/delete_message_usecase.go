package usecase

import (
	"hackathon/dao"
)

type DeleteMessageUseCase struct {
	MessageDao *dao.MessageDao
}

func (uc *DeleteMessageUseCase) DeleteMessage(id string) error {
	return uc.MessageDao.DeleteMessage(id)
}
