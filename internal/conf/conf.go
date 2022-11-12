package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Sever *ServerSettings `toml:"server"`
	DB    *MySQLSettingS  `toml:"mysql"`
	JWT   *JwtSettings    `toml:"jwt"`
}

var Conf *Config

func Init() (err error) {

	filePath := "internal/conf/config.toml"
	if _, err := toml.DecodeFile(filePath, &Conf); err != nil {
		return err
	}

	//查看配置
	fmt.Printf("这是 %v的设置:\n", Conf.DB)

	return nil

}
