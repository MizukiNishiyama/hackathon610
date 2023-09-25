package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"hackathon/model"
)

type ChannelDao struct {
	DB *sql.DB
}

func (dao *ChannelDao) FindByName(name string) ([]model.Channel, error) {
	rows, err := dao.DB.Query("SELECT channel_id, channel_name FROM channel WHERE channel_name = ?", name)
	if err != nil {
		return nil, err
	}

	channels := make([]model.Channel, 0)
	for rows.Next() {
		var u model.Channel
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			return nil, err
		}
		channels = append(channels, u)
	}

	return channels, nil
}

func (dao *ChannelDao) GetChannels() ([]model.Channel, error) {
	rows, err := dao.DB.Query("SELECT channel_id, channel_name FROM channel")
	if err != nil {
		return nil, err
	}

	channels := make([]model.Channel, 0)
	for rows.Next() {
		var u model.Channel
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			return nil, err
		}
		channels = append(channels, u)
	}

	return channels, nil
}

func (dao *ChannelDao) Insert(channel model.Channel) error {
	_, err := dao.DB.Exec("INSERT into channel VALUES(?, ?)", channel.Id, channel.Name)
	return err
}

func (dao *ChannelDao) DeleteChannel(id string) error {
	_, err := dao.DB.Exec("DELETE from channel WHERE channel_id =?", id)
	return err
}
