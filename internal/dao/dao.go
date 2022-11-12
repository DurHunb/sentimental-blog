package dao

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github/durhunb/emo-blog/internal/conf"
	"os"
)

// 未来加上redis等
type Dao struct {
	Db *sql.DB
}

func New(config *conf.Config) (d *Dao) {
	d = &Dao{
		Db: DBInit(config),
	}
	return
}

// DBInit 初始化mysql数据库
func DBInit(config *conf.Config) (Db *sql.DB) {
	//todo
	Db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		logrus.Errorf("Open mysql error:", err)
		os.Exit(-1)
	}
	return Db
}
