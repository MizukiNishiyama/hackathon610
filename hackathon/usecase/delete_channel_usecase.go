package usecase

import (
	"hackathon/dao"
)

type DeleteChannelUseCase struct {
	ChannelDao *dao.ChannelDao
}

func (uc *DeleteChannelUseCase) DeleteChannel(id string) error {
	return uc.ChannelDao.DeleteChannel(id)
}
