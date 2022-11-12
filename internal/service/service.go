package service

import (
	"github/durhunb/emo-blog/internal/conf"
	"github/durhunb/emo-blog/internal/dao"
)

type Service struct {
	config *conf.Config
	dao    *dao.Dao
}

// svc 服务的集合
var svc *Service

func Initialize(config *conf.Config) {

	svc = &Service{
		config: config,
		dao:    dao.New(config)}

	return
}
