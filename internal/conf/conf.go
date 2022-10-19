package conf

import "fmt"

var (
	ServerSetting *ServerSettings
	MySQLSetting  *MySQLSettingS
	JwtSetting    *JwtSettings
)

func setupSetting() error {
	viperSetting, err := InitSetting()

	if err != nil {
		return err
	}

	objects := map[string]interface{}{
		"Server": &ServerSetting,
		"MySQL":  &MySQLSetting,
		"Jwt":    &JwtSetting,
	}

	if err = viperSetting.Unmarshal(objects); err != nil {
		fmt.Printf("viper.Unmarshal(Conf) err %s", err)
	}

	//查看设置配置是否有误
	for name, setting := range objects {
		fmt.Printf("这是 %s的设置:%v\n", name, setting)
	}

	return nil
}

func Initialize() {
	err := setupSetting()
	if err != nil {
		return
	}

}
