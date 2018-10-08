package conf

import "github.com/spf13/viper"

//环境信息
const env = "dev"
const server = "localhost"

//基本配置接口
type Config struct {
	service string
	v       *viper.Viper
}

func (config *Config) LoadHarborConfig(configName string, configPath string) error {
	viper.Reset()
	viper.SetConfigFile("toml")
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	config.v = viper.GetViper()

	return nil
}

//返回指定key值的配置
func (config *Config) Get(key string) string {
	value := config.v.GetString(env + "." + config.service + "." + server + "." + key)
	if value == "" {
		value = config.v.GetString(env + "." + config.service + "." + key)
		if value == "" {
			value = config.v.GetString(env + "." + key)
			if value == "" {
				value = config.v.GetString(key)
			}
		}
	}
	return value
}
