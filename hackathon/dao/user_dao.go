package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"hackathon/model"
)

type UserDao struct {
	DB *sql.DB
}

func (dao *UserDao) FindByName(name string) ([]model.User, error) {
	rows, err := dao.DB.Query("SELECT id, name, email FROM user WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (dao *UserDao) Insert(user model.User) error {
	_, err := dao.DB.Exec("INSERT into user VALUES(?, ?, ?)", user.Id, user.Name, user.Email)
	return err
}
