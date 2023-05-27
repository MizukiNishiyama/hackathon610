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
	rows, err := dao.DB.Query("SELECT id, name FROM channel WHERE name = ?", name)
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
