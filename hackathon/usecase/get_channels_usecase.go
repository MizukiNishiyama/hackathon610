package usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

type GetChannelUseCase struct {
	ChannelDao *dao.ChannelDao
}

func (uc *GetChannelUseCase) Handle() ([]model.Channel, error) {
	return uc.ChannelDao.GetChannels()
}
