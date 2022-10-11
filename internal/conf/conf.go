package conf

import "fmt"

var (
	ServerSetting *ServerSettings
	MySQLSetting  *MySQLSettingS
)

func setupSetting() error {
	viperSetting, err := InitSetting()

	if err != nil {
		return err
	}

	objects := map[string]interface{}{
		"Server": &ServerSetting,
		"MySQL":  &MySQLSetting,
	}

	if err = viperSetting.Unmarshal(objects); err != nil {
		fmt.Printf("viper.Unmarshal(Conf) err %s", err)
	}

	fmt.Println("This is conf", objects)

	fmt.Println("this is serverSetting", ServerSetting)
	fmt.Println("this is MySQLSetting", MySQLSetting)

	return nil
}

func Initialize() {
	err := setupSetting()
	if err != nil {
		return
	}

}
