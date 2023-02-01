package dao

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github/durhunb/emo-blog/internal/model"
)

const (
	_getUserByUsernameSQL = "SElECT id,username,phone,`status`,`character` FROM user_account WHERE username = ? AND is_deleted = 0"
	_createUser           = "INSERT INTO user_account (nickname,username,password,encrypted_password,status) VALUES (?,?,?,?,?)"
)

func (d *AccountDao) GetUserByUsername(username string) (usr *model.User, err error) {
	rows, err := d.Db.Query(_getUserByUsernameSQL, username)
	if err != nil {
		logrus.Errorf("GetUserByUsername 查询用户名称 username(%s) error(%s)", username, err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("GetUserByUsername 关闭查询失败 error(%s)", err)
		}
	}(rows)

	for rows.Next() {
		usr := new(model.User)
		if err = rows.Scan(&usr.ID, &usr.Username, &usr.Phone, &usr.Status, &usr.Character); err != nil {
			logrus.Errorf("GetUserByUsername d.Db.Scan error=%s", err)
		}

	}
	return usr, nil
}

func (d *AccountDao) CreatUser(user *model.User) (err error) {
	_, err = d.Db.Exec(_createUser, user.Nickname, user.Username, user.Password, user.Encrypted_password, user.Status)
	if err != nil {
		logrus.Errorf("CreatUser 用户创建失败 err=%s", err)
		return err
	}
	return nil
}
