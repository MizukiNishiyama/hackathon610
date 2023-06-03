package usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

type SearchMessageUseCase struct {
	MessageDao *dao.MessageDao
}

func (uc *SearchMessageUseCase) Handle(channelid string) ([]model.Message, error) {
	return uc.MessageDao.ShowMessages(channelid)
}
