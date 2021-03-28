package confReading

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Db DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Connection string
	User       string
	Password   string
	Db         string
	Port       string
}

var spec *Configuration

func New() IConfigReader {
	return &Configuration{}
}

func (a Configuration) GetConf() *Configuration {
	return spec
}
func (a Configuration) LoadConfig() {
	viper.AddConfigPath("../config")
	viper.SetConfigName("conf.yml")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	// val := viper.GetString("app.datasource.host")
	err2 := viper.Unmarshal(&spec)
	if err2 != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

}
