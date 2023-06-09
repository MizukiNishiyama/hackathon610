package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"hackathon/model"
)

type MessageDao struct {
	DB *sql.DB
}

func (dao *MessageDao) ShowMessages(channelid string) ([]model.Message, error) {
	rows, err := dao.DB.Query("SELECT message_id, message_content, user_id, channel_id, time FROM message WHERE channel_id = ?", channelid)
	if err != nil {
		return nil, err
	}

	messages := make([]model.Message, 0)
	for rows.Next() {
		var u model.Message
		if err := rows.Scan(&u.Id, &u.Content, &u.UserId, &u.ChannelId, &u.Time); err != nil {
			return nil, err
		}
		messages = append(messages, u)
	}

	return messages, nil
}

func (dao *MessageDao) Insert(message model.Message) error {
	_, err := dao.DB.Exec("INSERT into message VALUES(?, ?, ?, ?, ?)", message.Id, message.Content, message.UserId, message.ChannelId, message.Time)
	return err
}

func (dao *MessageDao) DeleteMessage(id string) error {
	_, err := dao.DB.Exec("DELETE from message WHERE message_id =?", id)
	return err
}

func (dao *MessageDao) EditMessage(id, content string) error {
	_, err := dao.DB.Exec("UPDATE message SET message_content = CONCAT(message_content, '  （編集済み）') WHERE message_id = ?", id)
	return err
}
