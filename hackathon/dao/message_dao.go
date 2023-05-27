package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"hackathon/model"
)

type MessageDao struct {
	DB *sql.DB
}

func (dao *MessageDao) FindByName(content string) ([]model.Message, error) {
	rows, err := dao.DB.Query("SELECT id, content, userid, channelid FROM message WHERE content = ?", content)
	if err != nil {
		return nil, err
	}

	messages := make([]model.Message, 0)
	for rows.Next() {
		var u model.Message
		if err := rows.Scan(&u.Id, &u.Content, &u.UserId, &u.ChannelId); err != nil {
			return nil, err
		}
		messages = append(messages, u)
	}

	return messages, nil
}

func (dao *MessageDao) Insert(message model.Message) error {
	_, err := dao.DB.Exec("INSERT into message VALUES(?, ?, ?, ?)", message.Id, message.Content, message.UserId, message.ChannelId)
	return err
}
