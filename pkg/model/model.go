package model

import (
	"fmt"
	"github.com/leeeo2/backend/pkg/common/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Port        string `yaml:"Port"`
	Schema      string `yaml:"Schema"`
	MaxIdleConn int    `yaml:"MaxIdleConn"`
	MaxOpenConn int    `yaml:"MaxOpenConn"`
	Charset     string `yaml:"Charset"`
	Engine      string `yaml:"Engine"`
	Collate     string `yaml:"Collate"`
}

var tables = []interface{}{
	&User{},
}

var db *gorm.DB

func Setup(dbConf *Config, logConf *logger.Config, gormLogConf *logger.GormLogConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Schema, dbConf.Charset)
	options := fmt.Sprintf("ENGINE=%s DEFAULT CHARSET=%s COLLATE=%s", dbConf.Engine, dbConf.Charset, dbConf.Collate)

	logger, err := logger.NewGormLogger(logConf, gormLogConf)
	if err != nil {
		return err
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return err
	}
	err = db.Set("gorm:table_options", options).AutoMigrate(
		tables...,
	)
	if err != nil {
		return err
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(dbConf.MaxIdleConn)
	sqlDb.SetMaxOpenConns(dbConf.MaxOpenConn)
	return nil
}

func GetDB() *gorm.DB {
	return db
}
