package usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

type RegisterChannelUseCase struct {
	ChannelDao *dao.ChannelDao
}

func (uc *RegisterChannelUseCase) Handle(channel model.ChannelReqForHTTPPost) (model.Channel, error) {
	id := ulid.Make().String()
	channelToInsert := model.Channel{
		Id:   id,
		Name: channel.Name,
	}

	err := uc.ChannelDao.Insert(channelToInsert)
	if err != nil {
		return model.Channel{}, err
	}

	return channelToInsert, nil
}
