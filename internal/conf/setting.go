package conf

// 这里存放和配置文件映射的结构体

import (
	"time"
)

type ServerSettings struct {
	RunMode      string
	HttpIp       string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MySQLSettingS struct {
	UserName string
	Password string
	Host     string
	Port     int
	DBName   string
	//Charset      string
	//ParseTime    bool
	//MaxIdleConns int
	//MaxOpenConns int
}

type JwtSettings struct {
	//密钥
	Secret string
	//token签发者
	Issuer string
	//有效时间
	ExpireTime time.Duration
}
