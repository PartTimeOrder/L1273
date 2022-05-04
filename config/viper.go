package config

import (
	"L1273/model"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

// Viper 初始化配置文件
func Viper() *model.Config {
	c := &model.Config{}
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		return nil
	}
	if err := v.Unmarshal(&c); err != nil {
		return nil
	}
	return c
}
