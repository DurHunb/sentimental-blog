package conf

// 这里存放和配置文件映射的结构体

import (
	"github.com/spf13/viper"
	"time"
)

type ViperSetting struct {
	vp *viper.Viper
}

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

// InitSetting 初始化viper，配置toml文件
func InitSetting() (*ViperSetting, error) {

	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("toml")
	vp.AddConfigPath("internal/conf/")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &ViperSetting{vp}, nil
}

// Unmarshal 解组配置文件
func (s *ViperSetting) Unmarshal(objects map[string]interface{}) error {
	for k, v := range objects {
		err := s.vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
