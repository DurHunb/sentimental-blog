package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github/durhunb/emo-blog/internal/conf"
	"os"
)

// 未来加上redis等
type Dao struct {
	Db         *sql.DB
	AccountDao *AccountDao
}

type AccountDao struct {
	Db *sql.DB
}

func New(config *conf.Config) (d *Dao) {
	Db := DBInit(config)
	d = &Dao{
		Db:         Db,
		AccountDao: &AccountDao{Db: Db},
	}
	return
}

// DBInit 初始化mysql数据库
func DBInit(config *conf.Config) (Db *sql.DB) {
	//todo

	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DB.UserName, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName)

	Db, err := sql.Open("mysql", database)
	if err != nil {
		logrus.Errorf("Open mysql error:", err)
		os.Exit(-1)
	}
	return Db
}
