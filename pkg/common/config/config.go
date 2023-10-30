package config

import (
	"github.com/leeeo2/backend/pkg/common/logger"
	"github.com/leeeo2/backend/pkg/model"

	"github.com/spf13/viper"
)

type Config struct {
	Database model.Config         `yaml:"Database"`
	Server   ServerConfig         `yaml:"Server"`
	Log      logger.Config        `yaml:"Log"`
	GormLog  logger.GormLogConfig `yaml:"GormLog"`
}

type ServerConfig struct {
	ListenAddr string `yaml:"Host"`
	ListenPort int    `yaml:"Port"`
}

var globalConfig Config

func Global() *Config {
	return &globalConfig
}

func defaultConfig() {
	globalConfig = Config{
		Database: model.Config{
			User:        "root",
			Password:    "12345678",
			Host:        "127.0.0.1",
			Port:        "3306",
			Schema:      "backend",
			MaxIdleConn: 20,
			MaxOpenConn: 5,
			Charset:     "utf8",
			Engine:      "InnoDB",
			Collate:     "utf8_bin",
		},
		Server: ServerConfig{
			ListenAddr: "0.0.0.0",
			ListenPort: 8888,
		},
		Log: logger.Config{
			Filename:   "/var/log/backend/backend.log",
			MaxSize:    500,
			MaxAge:     7,
			MaxBackups: 30,
			LocalTime:  false,
			Compress:   false,
			CallerSkip: 3,
			Level:      "debug",
			Console:    "stdout",
		},
		GormLog: logger.GormLogConfig{
			GormLevel:                 "info",
			SqlSlowThreshold:          300,
			IgnoreRecordNotFoundError: false,
			IgnoreDuplicateError:      false,
		},
	}
}

func Init(configPath string) error {
	defaultConfig()
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&globalConfig)
	if err != nil {
		panic(err)
	}
	return nil
}
