package dao

import (
	"github.com/sirupsen/logrus"
	"github/durhunb/emo-blog/internal/model"
)

const (
	_getUserByUsernameSQL = "SElECT id,username,phone,`status`,`character` FROM user_account WHERE username = ? AND is_deleted = 0"
)

func (d *Dao) GetUserByUsername(username string) (res []*model.User, err error) {
	rows, err := d.Db.Query(_getUserByUsernameSQL, username)
	if err != nil {
		logrus.Errorf("GetUserByUsername d.Db.Query username(%s) error(%s)", username, err)
	}
	defer rows.Close()

	for rows.Next() {
		usr := new(model.User)
		if err = rows.Scan(&usr.ID, &usr.Username, &usr.Phone, &usr.Status, &usr.Character); err != nil {
			logrus.Errorf("etUserByUsername d.Db.Scan error(%s)", err)
		}
		res = append(res, usr)
	}
	return
}
