package usecase

import (
	"fmt"
	"hackathon/dao"
)

type EditMessageUseCase struct {
	MessageDao *dao.MessageDao
}

func (uc *EditMessageUseCase) EditMessage(id, content string) error {
	fmt.Println(id, content)
	return uc.MessageDao.EditMessage(id, content)
}
