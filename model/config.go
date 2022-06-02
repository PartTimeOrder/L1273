package model

import "time"

type Config struct {
	Api     string            `yaml:"api"`
	Address map[string]string `yaml:"address"`
	BotKey  string            `yaml:"botKey"`
	BotApi  string            `yaml:"botApi"`
	Num     time.Duration     `yaml:"num"`
}
