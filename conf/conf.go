package conf

import (
	"dbdoc/log"

	"github.com/spf13/viper"
)

var (
	Config    *viper.Viper
	Version   string
	BulidTime string
	GoVersion string
	Creater   string
)

/**
加载配置文件
*/
func InitConfig(configFile string) error {
	Config = viper.New()
	Config.SetConfigFile(configFile)
	err := Config.ReadInConfig()
	if err != nil {
		log.Println("InitConfig conf file:%v error", configFile)
		panic(err)
		return err
	}
	log.Println("InitConfig success: Version:%v BulidTime:%v GoVersion:%v Creater:%v env:%v", Version, BulidTime, GoVersion, Creater, Config.Get("env"))
	return nil
}
