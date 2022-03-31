package Config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var AppConf ConfigContainer

type ConfigContainer struct {
	Debug       bool
	Http_server string
	Http_port   int
	Mysql       MysqlConfig
}

type MysqlConfig struct {
	Host     string
	Port     int
	Prefix   string
	Username string
	Password string
	Charset  string
}

func NewConfigContainer() *ConfigContainer {
	c := &ConfigContainer{}
	c.GetAllConf()
	AppConf = *c
	return c
}

func (c *ConfigContainer) GetAllConf() {

	vjson := viper.New()
	vjson.SetConfigName("app")
	vjson.SetConfigType("json")
	vjson.AddConfigPath("conf")

	if err := vjson.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	vjson.Unmarshal(c)
	fmt.Println(time.Now().Format("01-01 00:00") + " 配置文件加载完毕!")

}
