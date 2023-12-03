package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/lib"
)

func init() {
	//取程序名称
	name := lib.AppName()

	//引入viper配置文件
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	//viper.SetEnvPrefix("database")

	_ = viper.BindPFlags(pflag.CommandLine)
}

func Load() error {
	return viper.ReadInConfig()
}

func Store() error {
	return viper.SafeWriteConfig()
}

func Register(module string, key string, value any) {
	viper.SetDefault(module+"."+key, value)
}

func GetBool(module string, key string) bool {
	return viper.GetBool(module + "." + key)
}

func GetString(module string, key string) string {
	return viper.GetString(module + "." + key)
}

func GetInt(module string, key string) int {
	return viper.GetInt(module + "." + key)
}

func GetFloat(module string, key string) float64 {
	return viper.GetFloat64(module + "." + key)
}
