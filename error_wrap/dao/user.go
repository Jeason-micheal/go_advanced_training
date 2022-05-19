package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"sql/dao/table"
)

func AddUser(u *table.User) (err error) {
	sqlStr := "INSERT INTO user(username, password, status, phone) VALUES(?, ?, ?, ?)"
	result, err := DB().Exec(sqlStr, u.Username, u.Password, u.Status, u.Phone)
	if err != nil {
		return errors.Wrap(err, "exit "+sqlStr+" failed")
	}
	_, err = result.LastInsertId()
	return errors.Wrap(err, "get insert lastInsertId failed")
}

func GetUserById(id int64) (user *table.User, err error) {
	sqlStr := "SELECT * FROM user WHERE id = ?"

	u1 := table.User{}
	err = DB().Get(&u1, sqlStr, id)
	//Sentinel Error
	if err == sql.ErrNoRows {
		err = errors.Wrap(err, fmt.Sprintf("Table user has no data id %d", id))
	} else {
		err = errors.Wrap(err, "Get failed")
	}
	user = &u1

	return
}

func GetAllUsers() (users []table.User, err error) {

	sqlStr := "SELECT * FROM user"
	err = DB().Select(&users, sqlStr)

	return
}

func UpdateUserNameById(userName string, id int64) (err error) {
	sqlStr := "UPDATE user SET username = ? WHERE id=?"
	result, err := DB().Exec(sqlStr, userName, id)
	if err != nil {
		return errors.Wrap(err, "Exec "+sqlStr+" failed")
	}
	_, err = result.RowsAffected()
	return errors.Wrap(err, "RowsAffected error")
}

func DeleteUserById(id int64) (err error) {
	sqlStr := "DELETE FROM user WHERE id = ?"
	result, err := DB().Exec(sqlStr, id)
	if err != nil {
		return errors.Wrap(err, "Exec "+sqlStr+" failed")
	}
	_, err = result.RowsAffected()
	return errors.Wrap(err, "RowsAffected failed")
}
