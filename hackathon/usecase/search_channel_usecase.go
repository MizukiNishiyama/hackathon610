package usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

type SearchChannelUseCase struct {
	ChannelDao *dao.ChannelDao
}

func (uc *SearchChannelUseCase) Handle(name string) ([]model.Channel, error) {
	return uc.ChannelDao.FindByName(name)
}
