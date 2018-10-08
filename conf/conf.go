// 配置接口，用于从配置文件读取配置信息
// 开发者:天河星光容器云开发小组(郭贵鑫、杜量、刘志聪、赵帅帅)
// 邮箱:guixin.guo@nscc-gz.cn、liang.du@nscc-gz.cn、zhicong.liu@nscc-gz.cn、shuaishuai.zhao@nscc-gz.cn
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

func (config *Config) LoadHarborConfig(service string) error {
	viper.Reset()
	viper.SetConfigFile("toml")
	viper.SetConfigName("nebula")
	viper.AddConfigPath("../../deployment/config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	config.v = viper.GetViper()
	config.service = service
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
