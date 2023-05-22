package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sourse/model"
)

type UserDao struct {
	DB *sql.DB
}

func (dao *UserDao) FindByName(name string) ([]model.User, error) {
	rows, err := dao.DB.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (dao *UserDao) Insert(user model.User) error {
	_, err := dao.DB.Exec("INSERT into user_2 VALUES(?, ?, ?)", user.Id, user.Name, user.Age)
	return err
}
