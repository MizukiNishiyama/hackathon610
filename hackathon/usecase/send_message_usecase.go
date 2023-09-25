package usecase

import (
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
)

type RegisterMessageUseCase struct {
	MessageDao *dao.MessageDao
}

func (uc *RegisterMessageUseCase) Handle(message model.MessageReqForHTTPPost) (model.Message, error) {
	id := ulid.Make().String()
	messageToInsert := model.Message{
		Id:        id,
		Content:   message.Content,
		UserId:    message.UserId,
		ChannelId: message.ChannelId,
		Time:      message.Time,
	}

	err := uc.MessageDao.Insert(messageToInsert)
	if err != nil {
		return model.Message{}, err
	}

	return messageToInsert, nil
}
